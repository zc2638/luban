/**
 * Created by zc on 2020/6/7.
 */
package global

import (
	"fmt"
	"github.com/mitchellh/go-homedir"
	"github.com/mitchellh/mapstructure"
	"github.com/spf13/viper"
	"github.com/zc2638/drone-control/global"
	"luban/pkg/database"
	"luban/pkg/server"
	"os"
	"strings"
)

type Config struct {
	Server   server.Config   `json:"server" yaml:"server"`
	Database database.Config `json:"database" yaml:"database"`
	Control  global.Config   `json:"control" yaml:"control"`
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
	cfg := Environ()
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(*os.PathError); ok {
			fmt.Println("Warning: not find config file.")
			return cfg, nil
		}
		return nil, err
	}
	fmt.Println("Using config file:", viper.ConfigFileUsed())
	err := viper.Unmarshal(cfg, func(dc *mapstructure.DecoderConfig) {
		dc.TagName = "json"
	})
	return cfg, err
}
