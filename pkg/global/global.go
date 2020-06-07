/**
 * Created by zc on 2020/6/7.
 */
package global

import "github.com/jinzhu/gorm"

func DB() *gorm.DB {
	return db
}
