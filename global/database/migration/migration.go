/**
 * Created by zc on 2020/6/3.
 */
package migration

import (
	"luban/global/database"
	"luban/pkg/store"
)

// auto migrate table struct
func InitTable(cfg *database.Config) error {
	db, err := database.New(cfg)
	if err != nil {
		return err
	}
	return db.AutoMigrate(
		&store.User{},
		&store.Space{},
		&store.Share{},
		&store.Resource{},
		&store.Version{},
		&store.Secret{},
		&store.Pipeline{},
	)
}
