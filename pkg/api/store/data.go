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
	Name string `json:"name"` // 名称
}

type Config struct {
	Name    string `json:"name" yaml:"name"`       // 名称
	Desc    string `json:"desc" yaml:"desc"`       // 描述
	Format  string `json:"format" yaml:"format"`   // 格式
	Content string `json:"content" yaml:"content"` // 内容
}

type ConfigVersion struct {
	Version string `json:"version"` // 版本号
	Format  string `json:"format"`  // 格式
}
