/**
* Author: Jason
* Email: jason_w96@163.com
* Date: 2023/10/8
* Time: 17:22
* Software: GoLand
 */

package elasticsearch

type Option struct {
	Addr       []string `json:"addr"`
	Username   string   `json:"username"`
	Password   string   `json:"password"`
	Index      string   `json:"index"`
	ScrollTime string   `json:"scroll_time"`
}
