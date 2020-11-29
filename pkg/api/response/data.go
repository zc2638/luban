/**
 * Created by zc on 2020/8/8.
**/
package response

type Timestamp struct {
	CreatedTS int64 `json:"created_ts"`
	UpdatedTS int64 `json:"updated_ts"`
}

type LoginResult struct {
	Username string `json:"username"`
	Token    string `json:"token"`
	Host     string `json:"host"`
}

type UserInfoResult struct {
	Username string `json:"username"`
	Host     string `json:"host"`
}

type SpaceResultItem struct {
	SpaceID string `json:"space_id"`
	Name    string `json:"name"`
	Timestamp
}

type ResourceResultItem struct {
	Name   string `json:"name"`
	Desc   string `json:"desc"`
	Kind   string `json:"kind"`
	Format string `json:"format"`
	Data   string `json:"data,omitempty"`
	Label  string `json:"label,omitempty"`
	Timestamp
}

type VersionResultItem struct {
	VersionID  string `json:"version_id"`
	ResourceID string `json:"resource_id"`
	Version    string `json:"version"`
	Kind       string `json:"kind"`
	Format     string `json:"format"`
	Desc       string `json:"desc"`
	Content    string `json:"content"`
	Timestamp
}

type TaskResult struct {
	TaskID     string               `json:"task_id"`
	PipelineID string               `json:"pipeline_id"`
	Status     int                  `json:"status"`
	StartTS    int64                `json:"start_ts"`
	EndTS      int64                `json:"end_ts"`
	CreatedTS  int64                `json:"created_ts"`
	UpdatedTS  int64                `json:"updated_ts"`
	Steps      []TaskStepResultItem `json:"steps"`
}

type TaskStepResultItem struct {
	StepID    string `json:"step_id"`
	Name      string `json:"name"`
	Status    int    `json:"status"`
	Log       string `json:"log"`
	StartTS   int64  `json:"start_ts"`
	EndTS     int64  `json:"end_ts"`
	CreatedTS int64  `json:"created_ts"`
	UpdatedTS int64  `json:"updated_ts"`
}
