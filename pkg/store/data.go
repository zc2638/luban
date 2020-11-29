/**
 * Created by zc on 2020/8/8.
**/
package store

import (
	"gorm.io/gorm"
	"time"
)

type DateTime struct {
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}

type User struct {
	UserID   string `json:"user_id" gorm:"primary_key"` // 唯一标识
	Username string `json:"username"`                   // 用户名
	Pwd      string `json:"pwd"`                        // 密码
	Salt     string `json:"salt"`                       // 盐值
	DateTime
}

type Space struct {
	SpaceID string `json:"space_id" gorm:"primary_key"` // 唯一标识
	UserID  string `json:"user_id"`                     // 用户标识
	Name    string `json:"name"`                        // 名称
	DateTime
}

type Share struct {
	ShareID string    `json:"share_id" gorm:"primary_key"` // 唯一标识
	SpaceID string    `json:"space_id"`                    // 空间标识
	UserID  string    `json:"user_id"`                     // 用户标识
	StartAt time.Time `json:"start_at"`                    // 开始时间
	EndAt   time.Time `json:"end_at"`                      // 结束时间
	DateTime
}

type Resource struct {
	ResourceID string `json:"resource_id" gorm:"primary_key"` // 唯一标识
	UserID     string `json:"user_id"`                        // 用户标识
	SpaceID    string `json:"space_id"`                       // 空间标识
	Name       string `json:"name"`                           // 名称
	Desc       string `json:"desc"`                           // 描述
	Kind       string `json:"kind"`                           // 类型(file、pipeline)
	Format     string `json:"format"`                         // 格式
	Data       string `json:"data"`                           // 内容
	Label      string `json:"label"`                          // 标签
	DateTime
}

type Version struct {
	VersionID  string `json:"version_id" gorm:"primary_key"` // 唯一标识
	ResourceID string `json:"resource_id"`                   // 资源标识
	Version    string `json:"version"`                       // 版本号
	Kind       string `json:"kind"`                          // 类型
	Format     string `json:"format"`                        // 格式
	Desc       string `json:"desc"`                          // 描述
	Data       string `json:"data"`                          // 内容
	DateTime
}

type Secret struct {
	SecretID string `json:"secret_id" gorm:"primary_key"` // 唯一标识
	UserID   string `json:"user_id"`                      // 用户标识
	SpaceID  string `json:"space_id"`                     // 空间标识
	Name     string `json:"name"`                         // 名称
	Value    string `json:"value"`                        // 值
	DateTime
}

type Pipeline struct {
	PipelineID string `json:"pipeline_id" gorm:"primary_key"` // 唯一标识
	UserID     string `json:"user_id"`                        // 用户标识
	SpaceID    string `json:"space_id"`                       // 空间标识
	ResourceID string `json:"config_id"`                      // 资源标识
	Name       string `json:"name"`                           // 名称
	Spec       string `json:"spec"`                           // 定义
	DateTime
}

const (
	TaskStatusPending int = iota
	TaskStatusRunning
	TaskStatusSuccess
	TaskStatusFail
)
