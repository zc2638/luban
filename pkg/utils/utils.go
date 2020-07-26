/**
 * Created by zc on 2020/7/25.
 */
package utils

func InStringSlice(ss []string, str string) (index int, exist bool) {
	for k, s := range ss {
		if s == str {
			return k, true
		}
	}
	return -1, false
}
