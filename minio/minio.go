/**
* Author: Jason
* Email: jason_w96@163.com
* Date: 2023/10/8
* Time: 17:49
* Software: GoLand
 */

package minio

import (
	"context"

	logger "github.com/LabKiko/kiko-logger"
	minioV7 "github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

const minioKey = "minio"

var Minio = new(minioConnector)

type minioConnector struct {
	clients map[string]minioImpl
}

func (minio *minioConnector) init(opts map[string]Option) error {
	var err error

	for key, option := range opts {
		cli := minioClient{opt: &option}
		cli.client, err = minio.initMinio(&option)
		if err != nil {
			logger.WithContext(context.Background()).WithError(err).WithFields(map[string]interface{}{"bucket": key}).Error("minio client init failed")
			return err
		}

		minio.clients[key] = &cli
	}

	return nil

}

func (minio *minioConnector) initMinio(opt *Option) (*minioV7.Client, error) {
	client, err := minioV7.New(opt.Address, &minioV7.Options{
		Creds:  credentials.NewStaticV4(opt.AccessKey, opt.SecretKey, ""),
		Secure: opt.SSL,
	})

	if err != nil {
		return nil, err
	}

	return client, nil
}

func (minio *minioConnector) Initializer() error {
	var minioConfig = new(Config)
	// 获取数据 todo

	if err := minio.init(minioConfig.Conf); err != nil {
		logger.WithContext(context.Background()).WithError(err).Error("minio init failed")
		return err
	}

	logger.Info("minio client init success")

	return nil
}

func (minio *minioConnector) Instance(bucket string) minioImpl {
	return minio.clients[bucket]
}

type minioImpl interface {
	Client() *minioV7.Client
	Option() *Option
}

type minioClient struct {
	opt    *Option
	client *minioV7.Client
}

func (minio *minioClient) Client() *minioV7.Client {
	return minio.client
}

func (minio *minioClient) Option() *Option {
	return minio.opt
}
