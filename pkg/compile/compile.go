/**
 * Created by zc on 2020/7/27.
 */
package compile

import "regexp"

func Name() *regexp.Regexp {
	return regexp.MustCompile("^[A-Za-z0-9][-A-Za-z0-9_.]*[A-Za-z0-9]$")
}
