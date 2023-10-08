/**
* Author: Jason
* Email: jason_w96@163.com
* Date: 2023/10/8
* Time: 15:40
* Software: GoLand
 */

package nats

type Options struct {
	Address  []string `json:"address"`
	Username string   `json:"username"`
	Password string   `json:"password"`
}
