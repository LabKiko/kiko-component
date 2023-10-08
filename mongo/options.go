/**
* Author: Jason
* Email: jason_w96@163.com
* Date: 2023/10/8
* Time: 15:40
* Software: GoLand
 */

package mongo

type Option struct {
	Address  string `json:"address"`
	Database string `json:"database"`
	Username string `json:"username"`
	Password string `json:"password"`
	Options  string `json:"options"`
	Direct   bool   `json:"direct"`
}

type MongoConfig struct {
	Conf map[string]Option
}
