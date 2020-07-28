/**
 * Created by zc on 2020/7/28.
 */
package compile

import "luban/pkg/errs"

var NameError = errs.Error("Invalid name, only support ^[A-Za-z0-9][-A-Za-z0-9_.]*[A-Za-z0-9]$")