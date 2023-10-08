/**
* Author: Jason
* Email: jason_w96@163.com
* Date: 2023/10/8
* Time: 17:22
* Software: GoLand
 */

package elasticsearch

import (
	"context"
	"fmt"
	"time"

	logger "github.com/LabKiko/kiko-logger"
	"github.com/olivere/elastic/v7"
)

const elasticSearchKey = "elastic_search"

var ElasticSearch = new(elasticSearchConnector)

type elasticSearchConnector struct {
	client *elastic.Client
	option *Option
}

func (es *elasticSearchConnector) Initializer() error {
	var opt = new(Option)
	// todo config 处理数据

	client, err := es.initElasticSearch(opt)
	if err != nil {
		logger.WithContext(context.TODO()).WithError(err).Error("elasticsearch init failed")
		return fmt.Errorf("elasticSearch init failed error: %s", err)
	}

	logger.Info("elasticsearch init success")

	es.client = client
	es.option = opt

	return nil

}

func (es *elasticSearchConnector) initElasticSearch(opt *Option) (*elastic.Client, error) {
	var options []elastic.ClientOptionFunc

	options = append(options, elastic.SetURL(opt.Addr...))
	options = append(options, elastic.SetSniff(false))
	options = append(options, elastic.SetBasicAuth(opt.Username, opt.Password))
	options = append(options, elastic.SetHealthcheckInterval(time.Second*10))

	client, err := elastic.NewClient(options...)
	if err != nil {
		logger.WithContext(context.Background()).WithError(err).Error("elasticsearch client init failed")
		return nil, fmt.Errorf("connect to elasticsearch failed, %v", err)
	}

	return client, nil

}

func (es *elasticSearchConnector) Instance() *elastic.Client {
	return es.client
}
func (es *elasticSearchConnector) Option() *Option {
	return es.option
}
