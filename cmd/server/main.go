/**
 * Created by zc on 2020/8/2.
 */
package main

import (
	"luban/cmd/server/app"
	"os"
)

func main() {
	command := app.NewServerCommand()
	if err := command.Execute(); err != nil {
		os.Exit(1)
	}
}
