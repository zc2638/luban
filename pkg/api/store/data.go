/**
 * Created by zc on 2020/6/4.
 */
package store

import "time"

type BaseTime struct {
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type BaseTimeWithDelete struct {
	BaseTime
	DeletedAt *time.Time `json:"deleted_at"`
}

type App struct {
	AppID string `json:"app_id"` // 应用标识
	SID   string `json:"sid"`    // 空间标识
	Title string `json:"title"`  // 标题
	Desc  string `json:"desc"`   // 描述
	BaseTimeWithDelete
}

type Config struct {
	ConfigID string `json:"config_id"` // 配置标识
	Title    string `json:"title"`     // 标题
	Content  string `json:"content"`   // 配置内容
	Version  string `json:"version"`   // 版本
	BaseTimeWithDelete
}

type ConfigRecord struct {
	ConfigID string    `json:"config_id"` // 配置标识
	Content  string    `json:"content"`   // 配置内容
	Version  string    `json:"version"`   // 版本
	Status   int       `json:"status"`    // 状态（0默认、1发布成功、2发布失败）
	CreateBy string    `json:"create_by"` // 创建用户标识
	CreateAt time.Time `json:"create_at"` // 创建时间
}

type ConfigRuleRelate struct {
	ConfigID    string `json:"config_id"`     // 配置标识
	SpaceRuleID string `json:"space_rule_id"` // 空间规则标识
}

type Rule struct {
	RuleID  string `json:"rule_id"` // 规则标识
	UID     string `json:"uid"`     // 用户标识
	Title   string `json:"title"`   // 标题
	Desc    string `json:"desc"`    // 描述
	Content string `json:"content"` // 规则内容
	Version string `json:"version"` // 版本
	BaseTimeWithDelete
}

type RuleRecord struct {
	RuleID   string    `json:"rule_id"`   // 规则标识
	Content  string    `json:"content"`   // 规则内容
	Version  string    `json:"version"`   // 版本
	CreateAt time.Time `json:"create_at"` // 创建时间
}

type Space struct {
	SID   string `json:"sid"`   // 空间标识
	Title string `json:"title"` // 标题
	Owner string `json:"owner"` // 所有人标识
	BaseTimeWithDelete
}

type SpaceAccess struct {
	SID    string `json:"sid"`    // 空间标识
	UID    string `json:"uid"`    // 用户标识
	Access string `json:"access"` // 空间权限集合
	BaseTime
}

type SpaceRule struct {
	SpaceRuleID string `json:"space_rule_id"` // 空间规则标识
	SID         string `json:"sid"`           // 空间标识
	Title       string `json:"title"`         // 标题
	Desc        string `json:"desc"`          // 描述
	Content     string `json:"content"`       // 规则内容
	CreateFrom  string `json:"create_from"`   // 创建来源用户标识
	BaseTimeWithDelete
}

type User struct {
	UID      string `json:"uid"`      // 用户标识
	Username string `json:"username"` // 用户名
	Email    string `json:"email"`    // 邮箱
	Avatar   string `json:"avatar"`   // 头像
	Pwd      string `json:"pwd"`      // 密码
	Salt     string `json:"salt"`     // 盐值
	BaseTimeWithDelete
}
