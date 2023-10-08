/**
* Author: Jason
* Email: jason_w96@163.com
* Date: 2023/10/8
* Time: 17:40
* Software: GoLand
 */

package nats

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"strings"
	"sync"
	"time"

	"github.com/nats-io/nats.go"
)

type Subscribe struct {
	Handler MessageHandler
	Entry   *nats.Subscription
}

type Client struct {
	conn       *nats.Conn
	subscribes map[string]*Subscribe

	lock *sync.Mutex
}

func NewClient(config *Options) (*Client, error) {
	client := &Client{
		lock:       &sync.Mutex{},
		subscribes: make(map[string]*Subscribe),
	}

	if err := client.Connect(config.Address...); err != nil {
		return nil, err
	}

	return client, nil
}

func (c *Client) Connect(servers ...string) (err error) {
	c.conn, err = nats.Connect(strings.Join(servers, ","),
		nats.Timeout(time.Second*10),
		nats.ErrorHandler(c.errorHandler),
	)

	if err != nil {
		return err
	}

	return nil
}

func (c *Client) Close() {
	for subject := range c.subscribes {
		c.UnSubscribe(subject)
	}

	if c.conn != nil {
		c.conn.Close()
	}
}

func (c *Client) Subscribe(subject string, handler MessageHandler) {
	c.lock.Lock()
	defer c.lock.Unlock()

	entry, err := c.conn.Subscribe(subject, c.dispatch)
	if err != nil {
		log.Fatal(err)
	}

	c.subscribes[subject] = &Subscribe{
		Handler: handler,
		Entry:   entry,
	}
}

func (c *Client) QueueSubscribe(subject string, queue string, handler MessageHandler) {
	c.lock.Lock()
	defer c.lock.Unlock()

	entry, err := c.conn.QueueSubscribe(subject, queue, c.dispatch)
	if err != nil {
		log.Fatal(err)
	}

	c.subscribes[subject] = &Subscribe{
		Handler: handler,
		Entry:   entry,
	}
}

func (c *Client) UnSubscribe(subject string) {
	c.lock.Lock()
	defer c.lock.Unlock()

	if subscribe, ok := c.subscribes[subject]; ok {
		_ = subscribe.Entry.Unsubscribe()
		delete(c.subscribes, subject)
	}
}

func (c *Client) Publish(subject string, data []byte) error {
	return c.conn.Publish(subject, data)
}

// 只有在建立连接后才能正常发送，重建连接的过程中不能发送
func (c *Client) PublishDirect(subject string, data []byte) error {
	if c.conn.Status() == nats.RECONNECTING {
		return nats.ErrConnectionReconnecting
	}
	return c.conn.Publish(subject, data)
}

func (c *Client) Request(subject string, data []byte, timeout time.Duration) (*Message, error) {
	msg, err := c.conn.Request(subject, data, timeout)
	if err != nil {
		return nil, err
	}

	return c.wrap(msg), nil
}

// 只有在建立连接后才能正常发送，重建连接的过程中不能发送
func (c *Client) RequestDirect(subject string, data []byte, timeout time.Duration) (*Message, error) {
	if c.conn.Status() == nats.RECONNECTING {
		return nil, nats.ErrConnectionReconnecting
	}
	msg, err := c.conn.Request(subject, data, timeout)
	if err != nil {
		return nil, err
	}

	return c.wrap(msg), nil
}

func (c *Client) dispatch(msg *nats.Msg) {
	if subscribe, ok := c.subscribes[msg.Subject]; ok {
		subscribe.Handler(c.wrap(msg))
	}
}

func (c *Client) errorHandler(nc *nats.Conn, s *nats.Subscription, err error) {
	if s != nil {
		log.Printf("async error in %q/%q: %v", s.Subject, s.Queue, err)
	} else {
		log.Printf("async error outside subscription: %v", err)
	}
}

func (c *Client) wrap(msg *nats.Msg) *Message {
	return &Message{
		conn:    c.conn,
		Subject: msg.Subject,
		Data:    msg.Data,
		Reply:   msg.Reply,
	}
}

type IEventbus interface {
	Connect(servers ...string) error
	Close()

	Subscribe(subject string, handler MessageHandler)
	UnSubscribe(subject string)

	QueueSubscribe(subject string, queue string, handler MessageHandler)

	Publish(subject string, data []byte) error
	Request(subject string, data []byte, timeout time.Duration) (*Message, error)

	PublishDirect(subject string, data []byte) error
	RequestDirect(subject string, data []byte, timeout time.Duration) (*Message, error)
}

type Message struct {
	conn *nats.Conn

	Subject string
	Data    []byte
	Reply   string
}

func (m *Message) Encode() ([]byte, error) {
	res, err := json.Marshal(m.Data)
	if err != nil {
		return nil, fmt.Errorf("message encode failed, %v", err)
	}

	return res, nil
}

func (m *Message) Decode(v interface{}) error {
	if err := json.Unmarshal(m.Data, v); err != nil {
		return fmt.Errorf("message decode failed, %v", err)
	}
	return nil
}

func (m *Message) Respond(data []byte) error {
	if m.Reply == "" {
		return ErrMsgNoReply
	}
	return m.conn.Publish(m.Reply, data)
}

type MessageHandler func(msg *Message)

var (
	ErrMsgNoReply = errors.New("event bus: message does not have a reply")
)
