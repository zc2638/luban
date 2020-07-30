/**
 * Created by zc on 2020/7/27.
 */
package compile

import "regexp"

var name = regexp.MustCompile(`^[A-Za-z0-9][-A-Za-z0-9_.]*[A-Za-z0-9]$`)

func Name() *regexp.Regexp {
	return name
}

var username = regexp.MustCompile(`^[a-zA-Z][a-z0-9A-Z]*([\-_.]([a-z0-9A-Z])+)*$`)

func Username() *regexp.Regexp {
	return username
}
