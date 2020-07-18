/**
 * Created by zc on 2020/6/7.
 */
package global

import (
	"luban/pkg/database"
	"luban/pkg/server"
)

type Config struct {
	Server   server.Config   `json:"server" yaml:"server"`
	Database database.Config `json:"database" yaml:"database"`
}

func Environ() *Config {
	// TODO 可以添加一些默认设置
	cfg := &Config{}
	cfg.Server.Secret = DefaultJwtSecret
	return cfg
}
