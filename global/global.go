/**
 * Created by zc on 2020/6/7.
 */
package global

import (
	"github.com/zc2638/drone-control/global"
	"gorm.io/gorm"
)

func DB() *gorm.DB {
	return db
}

func Cfg() *Config {
	return config
}

func ControlCfg() *global.Config {
	return global.Cfg()
}