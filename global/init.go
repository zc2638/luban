/**
 * Created by zc on 2020/6/7.
 */
package global

import (
	"github.com/jinzhu/gorm"
	"gopkg.in/yaml.v2"
	"luban/pkg/api/store"
	"luban/pkg/database"
	"luban/pkg/storage"
	"luban/pkg/uuid"
)

var config *Config

// InitConfig Initialize all used configurations
func InitConfig(cfg *Config) error {
	config = cfg
	if err := initDatabase(&cfg.Database); err != nil {
		return err
	}
	return initUser()
}

var db *gorm.DB

// InitDatabase Initialize database
func initDatabase(cfg *database.Config) error {
	var err error
	db, err = database.New(cfg)
	return err
}

func initUser() error {
	s := storage.New()
	data, err := s.Find(PathRoot, KeyUserFile)
	if err != nil {
		return err
	}
	if len(data) == 0 {
		users := []store.User{
			{
				Code:     uuid.New(),
				Username: "admin",
				Pwd:      "admin",
			},
		}
		b, err := yaml.Marshal(&users)
		if err != nil {
			return err
		}
		return s.Update(PathRoot, KeyUserFile, b)
	}
	return nil
}
