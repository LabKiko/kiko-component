/**
* Author: Jason
* Email: jason_w96@163.com
* Date: 2023/9/18
* Time: 15:53
* Software: GoLand
 */

package kafka

import (
	"github.com/segmentio/kafka-go"
)

const kafkaKey = "kafka"

type kafkaConnector struct {
	instances map[string]kafkaOpt
}

type kafkaOpt interface {
	Reader() *kafka.Reader
	Writer() *kafka.Writer
	Option() *Option
}

func (c *kafkaConnector) Instance(key string) {
}
