/**
 * Created by zc on 2020/6/7.
 */
package global

import (
	"github.com/zc2638/drone-control/global"
	"gorm.io/gorm"
	"luban/pkg/database"
	"net"
	"strconv"
)

var config *Config

// InitConfig Initialize all used configurations
func InitConfig(cfg *Config) error {
	config = cfg
	if err := InitControlConfig(cfg); err != nil {
		return err
	}
	return initDatabase(&cfg.Database)
}

func InitControlConfig(cfg *Config) error {
	controlConfig := &cfg.Control
	if controlConfig.RPC.Proto == "" {
		controlConfig.RPC.Proto = "http"
	}
	if controlConfig.RPC.Host == "" {
		controlConfig.RPC.Host = net.JoinHostPort("127.0.0.1", strconv.Itoa(cfg.Server.Port))
	}
	return global.InitCfg(controlConfig)
}

var db *gorm.DB

// InitDatabase Initialize database
func initDatabase(cfg *database.Config) error {
	var err error
	db, err = database.New(cfg)
	return err
}
