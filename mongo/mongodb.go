/**
* Author: Jason
* Email: jason_w96@163.com
* Date: 2023/9/18
* Time: 15:53
* Software: GoLand
 */

package mongo

import (
	"context"
	"fmt"

	logger "github.com/LabKiko/kiko-logger"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

const mongoKey = "mongo"

var MongoDB = new(mongoConnector)

/*
{
    "mongo":{
        "abc":{
            "address":""
        }
    }
}

*/

type mongoConnector struct {
	mongoClient map[string]*mongo.Client
}

func (c *mongoConnector) initMongo(config *Option) (*mongo.Client, error) {
	uri := fmt.Sprintf("mongodb://%s:%s@%s/%s%s", config.Username, config.Password, config.Address, config.Database, config.Options)
	opt := options.Client().ApplyURI(uri)

	if config.Direct == true {
		opt.SetDirect(true)
	}

	ctx := context.Background()
	client, err := mongo.Connect(ctx, opt)
	if err != nil {
		return nil, fmt.Errorf("connect to mongodb failed, %v", err)
	}

	if err = client.Ping(ctx, readpref.Primary()); err != nil {
		logger.WithContext(ctx).WithError(err).Error("mongo init failed")
		return nil, fmt.Errorf("ping mongodb failed, %v", err)
	}

	return client, nil
}

func (c *mongoConnector) Instance(db string) *mongo.Database {
	return c.mongoClient[db].Database(db)
}

func (c *mongoConnector) Initializer() error {
	var opt = new(MongoConfig)

	// todo  获取数据

	for key, value := range opt.Conf {
		client, err := c.initMongo(&value)
		if err != nil {
			logger.WithContext(context.Background()).WithFields(map[string]interface{}{"db": key}).Fatal("mongo init error")
			return fmt.Errorf("mongo db: %s init error: %s", key, err)
		}

		c.mongoClient[key] = client

	}

	logger.Info("mongo init success")

	return nil

}
