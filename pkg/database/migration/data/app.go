/**
 * Created by zc on 2020/6/4.
 */
package data

type App struct {
	AppID string `json:"app_id" gorm:"column:app_id;size:32;not null;primary_key;unique_index"` // 应用标识
	SID   string `json:"sid" gorm:"column:sid;size:32;not null"`                                // 空间标识
	Title string `json:"title" gorm:"not null"`                                                 // 标题
	Desc  string `json:"desc"`                                                                  // 描述
	BaseTimeWithDelete
}
