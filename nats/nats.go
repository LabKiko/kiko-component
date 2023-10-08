/**
* Author: Jason
* Email: jason_w96@163.com
* Date: 2023/9/18
* Time: 15:53
* Software: GoLand
 */

package nats

import (
	"context"

	logger "github.com/LabKiko/kiko-logger"
)

const natsKey = "nats"

var Nats = new(natsConnector)

type natsConnector struct {
	client *Client
}

func (nats *natsConnector) Initializer() error {
	var opt = new(Options)
	// TODO

	client, err := nats.initNats(opt)
	if err != nil {
		logger.WithContext(context.TODO()).WithError(err).Error("nats init failed")
		return err
	}

	nats.client = client

	logger.Info("nats init success")
	return nil
}

func (nats *natsConnector) initNats(opt *Options) (*Client, error) {

	client, err := NewClient(opt)
	if err != nil {
		logger.WithContext(context.Background()).WithError(err).Error("nats init failed")
		return nil, err
	}

	return client, nil
}

func (nats *natsConnector) Instance() *Client {
	return nats.client
}
