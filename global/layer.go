/**
 * Created by zc on 2020/6/7.
 */
package global

import (
	"github.com/pkgms/go/ctr"
	"github.com/sirupsen/logrus"
	"github.com/zc2638/drone-control/global"
	"gorm.io/gorm"
	"luban/global/database"
)

var config *Config

// InitConfig Initialize all used configurations
func InitConfig(cfg *Config) error {
	config = cfg
	logrus.SetFormatter(&logrus.TextFormatter{
		ForceColors:            true,
		DisableLevelTruncation: true,
		PadLevelText:           true,
		FullTimestamp:          true,
		TimestampFormat:        "2006/01/02 15:04:05",
	})
	ctr.InitLog(logrus.StandardLogger())
	if err := initControlConfig(&cfg.Control); err != nil {
		return err
	}
	return initDatabase(&cfg.Database)
}

func Cfg() *Config {
	return config
}

func initControlConfig(cfg *global.Config) error {
	return global.InitCfg(cfg)
}

func ControlCfg() *global.Config {
	return global.Cfg()
}

var db *gorm.DB

// InitDatabase Initialize database
func initDatabase(cfg *database.Config) error {
	var err error
	db, err = database.New(cfg)
	return err
}

func DB() *gorm.DB {
	return db
}
