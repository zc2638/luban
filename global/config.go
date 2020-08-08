/**
 * Created by zc on 2020/6/7.
 */
package global

import (
	"fmt"
	"github.com/mitchellh/go-homedir"
	"github.com/mitchellh/mapstructure"
	"github.com/spf13/viper"
	"luban/pkg/database"
	"luban/pkg/server"
	"strings"
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

func ParseConfig(cfgPath string) (*Config, error) {
	if cfgPath != "" {
		viper.SetConfigFile(cfgPath)
	} else {
		home, err := homedir.Dir()
		if err != nil {
			return nil, err
		}
		viper.AddConfigPath(home)
		viper.SetConfigName("config.yaml")
	}
	viper.SetEnvPrefix("LUBAN")
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv()
	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}
	fmt.Println("Using config file:", viper.ConfigFileUsed())
	cfg := Environ()
	err := viper.Unmarshal(cfg, func(dc *mapstructure.DecoderConfig) {
		dc.TagName = "json"
	})
	return cfg, err
}