/**
 * Created by zc on 2020/6/7.
 */
package global

import "gorm.io/gorm"

func DB() *gorm.DB {
	return db
}

func Cfg() *Config {
	return config
}
