/**
 * Created by zc on 2020/9/1.
 */
package database

import (
	"errors"
	"gorm.io/gorm"
)

func RecordNotFound(err error) bool {
	return errors.Is(err, gorm.ErrRecordNotFound)
}