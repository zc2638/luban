/**
 * Created by zc on 2020/8/2.
 */
package data

import (
	"sync"
)

type Queue struct {
	mux      sync.Mutex
	Waiting  []Task
	Running  []Task
	Complete []Task
}

type Task struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	User        string `json:"user"`
	Space       string `json:"space"`
	Config      string `json:"config"`
	Version     string `json:"version"`
	CreatedAt   int64  `json:"created_at"`
	StartedAt   int64  `json:"started_at"`
	CanceledAt  int64  `json:"canceled_at"`
	CompletedAt int64  `json:"completed_at"`
}
