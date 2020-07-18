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
}

type SpaceParams struct {
	Name string `json:"name"`
}
