/**
 * Created by zc on 2020/6/3.
 */
package migration

import (
	"errors"
	"luban/pkg/database"
	"luban/pkg/database/data"
)

// check database.
// create database if not exists
func InitDatabase(cfg *database.Config) error {
	if cfg.DBName == "" {
		return errors.New("dbname is empty")
	}
	dbConfig := cfg.Clone()
	dbConfig.DBName = ""
	db, err := database.New(dbConfig)
	if err != nil {
		return err
	}
	return db.Exec("CREATE DATABASE IF NOT EXISTS " + cfg.DBName + " DEFAULT CHARSET utf8mb4 COLLATE utf8mb4_general_ci").Error
}

// auto migrate table struct
func InitTable(cfg *database.Config) error {
	db, err := database.New(cfg)
	if err != nil {
		return err
	}
	return db.AutoMigrate(
		&data.User{},
		&data.Space{},
		&data.Share{},
		&data.Resource{},
		&data.Version{},
		&data.Secret{},
		&data.Pipeline{},
		&data.Task{},
		&data.TaskStep{},
	)
}
