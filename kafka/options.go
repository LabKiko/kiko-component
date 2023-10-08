/**
* Author: Jason
* Email: jason_w96@163.com
* Date: 2023/10/8
* Time: 15:40
* Software: GoLand
 */

package kafka

type Option struct {
	Address  []string `json:"address"`
	Username string   `json:"username"`
	Password string   `json:"password"`
	Topic    string   `json:"topic"`
	Group    string   `json:"group"`
}

type Config struct {
	Conf map[string]Option
}
