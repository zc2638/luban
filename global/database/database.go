/**
 * Created by zc on 2020/6/3.
 */
package database

import (
	"github.com/go-sql-driver/mysql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"time"
)

type Config struct {
	Addr     string `json:"addr" yaml:"addr"`
	Username string `json:"username" yaml:"username"`
	Password string `json:"password" yaml:"password"`
	DBName   string `json:"dbname" yaml:"dbname"`
	Debug    bool   `json:"debug" yaml:"debug"`
}

func (c *Config) Clone() *Config {
	return &Config{
		Addr:     c.Addr,
		Username: c.Username,
		Password: c.Password,
		DBName:   c.DBName,
		Debug:    c.Debug,
	}
}

func New(cfg *Config) (*gorm.DB, error) {
	// TODO select driver by config
	cfg.Debug = true
	db, err := gorm.Open(
		sqlite.Open("luban.db"),
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
	return db, nil
}

func buildMysqlDsn(cfg *Config) string {
	config := mysql.Config{
		Addr:                 cfg.Addr,
		User:                 cfg.Username,
		Passwd:               cfg.Password,
		DBName:               cfg.DBName,
		Collation:            "utf8mb4_general_ci",
		ParseTime:            true,
		Loc:                  time.UTC,
		AllowNativePasswords: true,
	}
	return config.FormatDSN()
}
