/**
 * Created by zc on 2020/6/9.
 */
package request

type RegisterParams struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginParams struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type SpaceParams struct {
	Name string `json:"name"`
}

type ResourceParams struct {
	Name    string `json:"name"`    // 名称
	Desc    string `json:"desc"`    // 描述
	Format  string `json:"format"`  // 格式
	Content string `json:"content"` // 内容
	Version string `json:"version"` // 版本
	Label   string `json:"label"`   // 标签
}

type ResourceVersionParams struct {
	Version string `json:"version"`
	Desc    string `json:"desc"`
}

type PipelineParams struct {
	ResourceID string `json:"resource_id"`
	Name       string `json:"name"`
	Spec       string `json:"spec"`
}

type TaskStepParams struct {
	Status int    `json:"status"`
	Log    string `json:"log"`
}
