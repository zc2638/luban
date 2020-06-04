/**
 * Created by zc on 2020/6/4.
 */
package data

import "time"

type BaseTime struct {
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type BaseTimeWithDelete struct {
	BaseTime
	DeletedAt *time.Time `json:"deleted_at" sql:"index"`
}
