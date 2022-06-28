/**
 * Created by zc on 2020/6/3.
 */
package database

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"luban/pkg/store"
	"time"
)

type Config struct {
	Driver     string `json:"driver"`
	Datasource string `json:"datasource"`
	Debug      bool   `json:"debug"`
}

func New(cfg *Config) (*gorm.DB, error) {
	var dialector gorm.Dialector
	if cfg.Driver == "mysql" {
		dialector = mysql.Open(cfg.Datasource)
	} else {
		if cfg.Datasource == "" {
			cfg.Datasource = "luban.db"
		}
		dialector = sqlite.Open(cfg.Datasource)
	}
	db, err := gorm.Open(
		dialector,
		&gorm.Config{
			SkipDefaultTransaction: true,
			NamingStrategy: schema.NamingStrategy{
				SingularTable: true,
			},
		},
	)
	if err != nil {
		return nil, err
	}
	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}
	sqlDB.SetMaxIdleConns(10)               // 最大空闲连接数
	sqlDB.SetMaxOpenConns(100)              // 最大打开连接数
	sqlDB.SetConnMaxLifetime(time.Hour * 6) // 连接最长存活时间
	if cfg.Debug {
		db = db.Debug()
	}
	if err := db.AutoMigrate(
		&store.User{},
		&store.Space{},
		&store.Share{},
		&store.Resource{},
		&store.Version{},
		&store.Secret{},
		&store.Pipeline{},
	); err != nil {
		fmt.Println("migration table failed")
		return nil, err
	}
	return db, nil
}
