/**
* Author: Jason
* Email: jason_w96@163.com
* Date: 2023/8/25
* Time: 12:00
* Software: GoLand
 */

package redis

import (
	"context"
	"fmt"
	"strconv"

	logger "github.com/LabKiko/kiko-logger"
	redisV9 "github.com/redis/go-redis/v9"
)

const redisKey = "redis"

var Redis = new(redisConnector)

type redisConnector struct {
	client redisV9.UniversalClient
}

func (c *redisConnector) Initializer() error {
	var opt = new(UniversalOption)
	// config 获取数据 TODO

	client, err := c.initRedis(opt)
	if err != nil {
		return fmt.Errorf("init redis had an error: %s", err)
	}

	logger.Info("redis init success")

	c.client = client

	return nil
}

func (c *redisConnector) initRedis(opt *UniversalOption) (redisV9.UniversalClient, error) {
	db, _ := strconv.Atoi(strconv.Itoa(opt.DB))
	poolSize, _ := strconv.Atoi(strconv.Itoa(opt.PoolSize))
	maxTries, _ := strconv.Atoi(strconv.Itoa(opt.MaxRetries))

	if poolSize == 0 {
		poolSize = 5
	}

	if maxTries == 0 {
		maxTries = 3
	}
	redisOpt := &redisV9.UniversalOptions{
		Addrs:      opt.Address,
		Password:   opt.Password,
		DB:         db,
		PoolSize:   poolSize,
		MaxRetries: maxTries,
		MasterName: opt.Master,
		ReadOnly:   opt.ReadOnly,
	}

	if opt.Random {
		redisOpt.RouteRandomly = true
	} else {
		redisOpt.RouteByLatency = true
	}

	client := redisV9.NewUniversalClient(redisOpt)
	ctx := context.Background()
	if res, err := client.Ping(context.Background()).Result(); err != nil {
		logger.WithContext(context.TODO()).WithFields(map[string]interface{}{
			"redis":  client.Info(ctx),
			"result": res,
		}).Error("redis init")
		return nil, fmt.Errorf("redis init ping err: %v", err)
	}

	return client, nil

}

func (c *redisConnector) Instance() redisV9.UniversalClient {
	return c.client
}
