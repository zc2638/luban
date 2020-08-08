/**
 * Created by zc on 2020/8/2.
 */
package env

import (
	"luban/global"
	"os"
)

func Config() string {
	cfgFileENV := os.Getenv("LUBAN_CONFIG")
	if cfgFileENV != "" {
		return cfgFileENV
	}
	return global.DefaultPath
}
