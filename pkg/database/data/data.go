/**
 * Created by zc on 2020/8/8.
**/
package data

import (
	"gorm.io/gorm"
	"time"
)

type DateTime struct {
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
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
	SpaceID    string `json:"space_id"`                       // 空间标识
	Name       string `json:"name"`                           // 名称
	Desc       string `json:"desc"`                           // 描述
	Format     string `json:"format"`                         // 格式
	Content    string `json:"content"`                        // 内容
	Label      string `json:"label"`                          // 标签
	DateTime
}

type Version struct {
	VersionID  string `json:"version_id" gorm:"primary_key"` // 唯一标识
	ResourceID string `json:"resource_id"`                   // 资源标识
	Version    string `json:"version"`                       // 版本号
	Format     string `json:"format"`                        // 格式
	Desc       string `json:"desc"`                          // 描述
	Content    string `json:"content"`                       // 内容
	DateTime
}

type Secret struct {
	SecretID string `json:"secret_id" gorm:"primary_key"` // 唯一标识
	SpaceID  string `json:"space_id"`                     // 空间标识
	Name     string `json:"name"`                         // 名称
	Value    string `json:"value"`                        // 值
	DateTime
}

type Pipeline struct {
	PipelineID string `json:"pipeline_id" gorm:"primary_key"` // 唯一标识
	SpaceID    string `json:"space_id"`                       // 空间标识
	ResourceID string `json:"config_id"`                      // 资源标识
	Name       string `json:"name"`                           // 名称
	Spec       string `json:"spec"`                           // 定义
	DateTime
}

type Task struct {
	TaskID     string    `json:"task_id"`     // 唯一标识
	PipelineID string    `json:"pipeline_id"` // 流水线标识
	Spec       string    `json:"content"`     // 流水线定义
	Data       string    `json:"data"`        // 资源内容
	Status     int       `json:"status"`      // 状态
	StartAt    time.Time `json:"start_at"`    // 开始时间
	EndAt      time.Time `json:"end_at"`      // 结束时间
	DateTime
}

type TaskStep struct {
	StepID  string    `json:"step_id"`  // 唯一标识
	TaskID  string    `json:"task_id"`  // 任务标识
	Name    string    `json:"name"`     // 名称
	Status  int       `json:"status"`   // 状态
	Log     string    `json:"log"`      // 日志
	StartAt time.Time `json:"start_at"` // 开始时间
	EndAt   time.Time `json:"end_at"`   // 结束时间
	DateTime
}

const (
	TaskStatusPending int = iota
	TaskStatusRunning
	TaskStatusSuccess
	TaskStatusFail
)
