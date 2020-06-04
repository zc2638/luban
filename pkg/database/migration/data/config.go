/**
 * Created by zc on 2020/6/4.
 */
package data

import "time"

type Config struct {
	ConfigID string `json:"config_id" gorm:"column:config_id;size:32;not null;primary_key;unique_index"` // 配置标识
	Title    string `json:"title" gorm:"not null"`                                                       // 标题
	Content  string `json:"content" gorm:"type:text;not null"`                                           // 配置内容
	Version  string `json:"version" gorm:"size:20;not null"`                                             // 版本
	BaseTimeWithDelete
}

type ConfigRecord struct {
	ConfigID string    `json:"config_id" gorm:"column:config_id;size:32;not null;primary_key"` // 配置标识
	Content  string    `json:"content" gorm:"type:text;not null"`                              // 配置内容
	Version  string    `json:"version" gorm:"size:20;not null"`                                // 版本
	Status   int       `json:"status" gorm:"type:tinyint"`                                     // 状态（0默认、1发布成功、2发布失败）
	CreateBy string    `json:"create_by" gorm:"size:32;not null"`                              // 创建用户标识
	CreateAt time.Time `json:"create_at"`                                                      // 创建时间
}

type ConfigRuleRelate struct {
	ConfigID    string `json:"config_id" gorm:"column:config_id;size:32;not null;primary_key"`         // 配置标识
	SpaceRuleID string `json:"space_rule_id" gorm:"column:space_rule_id;size:32;not null;primary_key"` // 空间规则标识
}
