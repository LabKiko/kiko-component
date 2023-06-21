/**
* Author: Jason
* Email: jason_w96@163.com
* Date: 2023/8/25
* Time: 11:58
* Software: GoLand
 */

package elasticsearch

var defaultURL = "http://localhost:9200"

type EsOptions struct {
	Address  []string `json:"address"`
	Username string   `json:"username"`
	Password string   `json:"password"`
}
type Option func(opts *EsOptions)

func NewEsOpts(opts ...Option) *EsOptions {
	esOp := &EsOptions{}
	for _, opt := range opts {
		opt(esOp)
	}

	return esOp
}

func WithAddress(address ...string) Option {
	return func(opts *EsOptions) {
		opts.Address = append(opts.Address, address...)
	}
}

func WithUserName(user string) Option {
	return func(opts *EsOptions) {
		opts.Username = user
	}
}

func WithPassword(password string) Option {
	return func(opts *EsOptions) {
		opts.Password = password
	}
}
