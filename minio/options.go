/**
* Author: Jason
* Email: jason_w96@163.com
* Date: 2023/10/8
* Time: 17:49
* Software: GoLand
 */

package minio

type Option struct {
	Address     string `json:"address"`
	Bucket      string `json:"bucket"`
	AccessKey   string `json:"access_key"`
	SecretKey   string `json:"secret_key"`
	SSL         bool   `json:"ssl"`
	DownloadURL string `json:"download_url"`
}

type Config struct {
	Conf map[string]Option
}
