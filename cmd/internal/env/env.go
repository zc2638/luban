/**
 * Created by zc on 2020/8/2.
 */
package env

import (
	"os"
)

func Config() string {
	return os.Getenv("LUBAN_CONFIG")
}
