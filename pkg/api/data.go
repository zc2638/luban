/**
 * Created by zc on 2020/6/9.
 */
package api

type RegisterParams struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginParams struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginResult struct {
	Username string `json:"username"`
	Token    string `json:"token"`
	Host     string `json:"host"`
}

type SpaceParams struct {
	Name string `json:"name"`
}

type ConfigParams struct {
	Name    string `json:"name"`    // 名称
	Desc    string `json:"desc"`    // 描述
	Format  string `json:"format"`  // 格式
	Content string `json:"content"` // 内容
	Version string `json:"version"` // 版本
}

type ConfigVersionParams struct {
	Version string `json:"version"` // 版本号
}
