//go:build !test

package main

import (
	"github.com/gookit/color"
	"github.com/woodpecker-kit/woodpecker-plugin-env"
	"github.com/woodpecker-kit/woodpecker-plugin-env/cmd/cli"
	"github.com/woodpecker-kit/woodpecker-plugin-env/internal/pkgJson"
	"github.com/woodpecker-kit/woodpecker-tools/wd_log"
	os "os"
)

func main() {
	pkgJson.InitPkgJsonContent(woodpecker_plugin_env.PackageJson)

	wd_log.SetLogLineDeep(wd_log.DefaultExtLogLineMaxDeep)

	app := cli.NewCliApp()

	args := os.Args
	if err := app.Run(args); nil != err {
		color.Redf("cli err at %v\n", err)
	}
}
