/**
 * Created by zc on 2020/6/4.
 */
package data

import "time"

type Rule struct {
	RuleID  string `json:"rule_id" gorm:"column:rule_id;size:32;not null;primary_key;unique_index"` // 规则标识
	UID     string `json:"uid" gorm:"column:uid;size:32;not null"`                                  // 用户标识
	Title   string `json:"title" gorm:"not null"`                                                   // 标题
	Desc    string `json:"desc"`                                                                    // 描述
	Content string `json:"content" gorm:"type:text;not null"`                                       // 规则内容
	Version string `json:"version" gorm:"size:20;not null"`                                         // 版本
	BaseTimeWithDelete
}

type RuleRecord struct {
	RuleID   string    `json:"rule_id" gorm:"column:rule_id;size:32;not null;primary_key"` // 规则标识
	Content  string    `json:"content" gorm:"type:text;not null"`                          // 规则内容
	Version  string    `json:"version" gorm:"size:20;not null"`                            // 版本
	CreateAt time.Time `json:"create_at"`                                                  // 创建时间
}
