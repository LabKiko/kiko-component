/**
* Author: Jason
* Email: jason_w96@163.com
* Date: 2023/8/25
* Time: 11:58
* Software: GoLand
 */

package redis

import (
	"github.com/redis/go-redis/v9"
)

type Opt func(o *Option)

type ClusterOpt func(o *ClusterOption)

type UniversalOpt func(o *UniversalOption)
type Option struct {
	// Mode       string   `json:"mode"` // sample ,cluster,
	Address    []string `json:"address"`
	Password   string   `json:"password"`
	User       string   `json:"user"`
	DB         int      `json:"db"`
	PoolSize   int      `json:"pool_size"`
	MaxRetries int      `json:"max_retries"`
}

func WithAddress(address ...string) Opt {
	return func(o *Option) {
		o.Address = append(o.Address, address...)
	}
}

func WithPassword(password string) Opt {
	return func(o *Option) {
		o.Password = password
	}
}

func WithDB(db int) Opt {
	return func(o *Option) {
		o.DB = db
	}
}

func WithUser(user string) Opt {
	return func(o *Option) {
		o.User = user
	}
}

type ClusterOption struct {
	Option
	Random   bool                `json:"random"`
	ReadOnly bool                `json:"read_only"`
	Nodes    []redis.ClusterNode `json:"nodes"`
}

func WithNodes(nodes ...redis.ClusterNode) ClusterOpt {
	return func(o *ClusterOption) {
		o.Nodes = append(o.Nodes, nodes...)
	}
}
func WithRandom(isRandom bool) ClusterOpt {
	return func(o *ClusterOption) {
		o.Random = isRandom
	}
}

func WithClusterAddress(address ...string) ClusterOpt {
	return func(o *ClusterOption) {
		o.Address = append(o.Address, address...)
	}
}

func WithClusterPassword(password string) ClusterOpt {
	return func(o *ClusterOption) {
		o.Password = password
	}
}

func WithClusterDB(db int) ClusterOpt {
	return func(o *ClusterOption) {
		o.DB = db
	}
}

func WithClusterUser(user string) ClusterOpt {
	return func(o *ClusterOption) {
		o.User = user
	}
}

type UniversalOption struct {
	ClusterOption
	Master string `json:"master"`
}

func WithMaster(master string) UniversalOpt {
	return func(o *UniversalOption) {
		o.Master = master
	}
}

func WithUniversalAddress(address ...string) UniversalOpt {
	return func(o *UniversalOption) {
		o.Address = append(o.Address, address...)
	}
}
