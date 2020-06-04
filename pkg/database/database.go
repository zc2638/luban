/**
 * Created by zc on 2020/6/3.
 */
package database

import (
	"github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"time"
)

type Config struct {
	Addr     string `json:"addr"`
	Username string `json:"username"`
	Password string `json:"password"`
	DBName   string `json:"dbname"`
	Debug    bool   `json:"debug"`
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

func NewDatabase(cfg *Config) (*gorm.DB, error) {
	dsn := buildMysqlDsn(cfg)
	db, err := gorm.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}

	db.DB().SetMaxIdleConns(10)               // 连接池的空闲数大小
	db.DB().SetMaxOpenConns(100)              // 最大打开连接数
	db.DB().SetConnMaxLifetime(time.Hour * 6) // 连接最长存活时间
	db.SingularTable(true)                    // 禁用复数表名
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
