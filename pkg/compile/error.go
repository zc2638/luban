/**
 * Created by zc on 2020/7/28.
 */
package compile

import "luban/pkg/errs"

const (
	NameError = errs.Error("Invalid name, only support ^[A-Za-z0-9][-A-Za-z0-9_.]*[A-Za-z0-9]$")
	UsernameError = errs.Error("Invalid username, only support ^[a-zA-Z][a-z0-9A-Z]*([\\-_.]([a-z0-9A-Z])+)*$")
)
