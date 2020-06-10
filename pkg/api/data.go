/**
 * Created by zc on 2020/6/9.
 */
package api

type RegisterParams struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginParams struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginResult struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Token    string `json:"token"`
}
