/**
 * Created by zc on 2020/6/4.
 */
package store

type User struct {
	Code     string `json:"code" yaml:"code"`         // 用户标识
	Username string `json:"username" yaml:"username"` // 用户名
	Pwd      string `json:"pwd" yaml:"pwd"`           // 密码
}

type Space struct {
	Name string `json:"name" yaml:"name"` // 名称
}

type Config struct {
	Name    string `json:"name" yaml:"name"`       // 名称
	Desc    string `json:"desc" yaml:"desc"`       // 描述
	Format  string `json:"format" yaml:"format"`   // 格式
	Content string `json:"content" yaml:"content"` // 内容
	Label   string `json:"label" yaml:"label"`     // 标签
}

type ConfigVersion struct {
	Version   string `json:"version" yaml:"version"`       // 版本号
	Format    string `json:"format" yaml:"format"`         // 格式
	Desc      string `json:"desc" yaml:"desc"`             // 描述
	IsLatest  bool   `json:"is_latest" yaml:"is_latest"`   // 是否默认
	CreatedAt int64  `json:"created_at" yaml:"created_at"` // 创建时间
}

type ConfigVersionGroup []ConfigVersion

func (cv ConfigVersionGroup) Len() int {
	return len(cv)
}
func (cv ConfigVersionGroup) Swap(i, j int) {
	cv[i], cv[j] = cv[j], cv[i]
}
func (cv ConfigVersionGroup) Less(i, j int) bool {
	return cv[j].CreatedAt < cv[i].CreatedAt
}
