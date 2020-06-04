/**
 * Created by zc on 2020/6/3.
 */
package data

type User struct {
	UID      string `json:"uid" gorm:"column:uid;size:32;not null;primary_key;unique_index"` // 用户标识
	Nickname string `json:"nickname" gorm:"column:nickname;size:50;not null"`                // 昵称
	Email    string `json:"email" gorm:"size:100"`                                           // 邮箱
	Pwd      string `json:"pwd" gorm:"type:varchar(100);not null"`                           // 密码
	Salt     string `json:"salt" gorm:"size:10;not null"`                                    // 盐值
	BaseTimeWithDelete
}
