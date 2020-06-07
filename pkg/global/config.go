/**
 * Created by zc on 2020/6/7.
 */
package global

import (
	"stone/pkg/database"
	"stone/pkg/serve"
)

type Config struct {
	Serve    serve.Config    `json:"serve"`
	Database database.Config `json:"database"`
}

func Environ() *Config {
	// TODO 可以添加一些默认设置
	return &Config{}
}
