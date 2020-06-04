/**
 * Created by zc on 2020/6/4.
 */
package config

import "stone/pkg/database"

const DefaultPath = "config.yaml"

type Config struct {
	Database database.Config `json:"database"`
}

func Environ() *Config {
	// TODO 可以添加一些默认设置
	return &Config{}
}
