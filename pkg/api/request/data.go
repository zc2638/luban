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
	Version string `json:"version"` // 版本号
	Desc    string `json:"desc"`    // 描述
}
