/**
 * Created by zc on 2020/6/4.
 */
package data

type Space struct {
	SID   string `json:"sid" gorm:"column:sid;size:32;not null;primary_key;unique_index"` // 空间标识
	Title string `json:"title" gorm:"not null"`                                           // 标题
	Owner string `json:"owner" gorm:"size:32;not null"`                                   // 所有人标识
	BaseTimeWithDelete
}

type SpaceAccess struct {
	SID    string `json:"sid" gorm:"column:sid;size:32;not null;primary_key"` // 空间标识
	UID    string `json:"uid" gorm:"column:uid;size:32;not null"`             // 用户标识
	Access string `json:"access"`                                             // 空间权限集合
	BaseTime
}

type SpaceRule struct {
	SpaceRuleID string `json:"space_rule_id" gorm:"column:space_rule_id;size:32;not null;primary_key;unique_index"` // 空间规则标识
	SID         string `json:"sid" gorm:"column:sid;size:32;not null;primary_key"`                                  // 空间标识
	Title       string `json:"title" gorm:"not null"`                                                               // 标题
	Desc        string `json:"desc"`                                                                                // 描述
	Content     string `json:"content" gorm:"type:text;not null"`                                                   // 规则内容
	CreateFrom  string `json:"create_from" gorm:"size:32;not null"`                                                 // 创建来源用户标识
	BaseTimeWithDelete
}
