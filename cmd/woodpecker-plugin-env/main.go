//go:build !test

package main

import (
	"github.com/gookit/color"
	"github.com/joho/godotenv"
	"github.com/woodpecker-kit/woodpecker-plugin-env"
	"github.com/woodpecker-kit/woodpecker-plugin-env/cmd/cli"
	"github.com/woodpecker-kit/woodpecker-plugin-env/internal/pkgJson"
	"github.com/woodpecker-kit/woodpecker-tools/wd_log"
	os "os"
)

func main() {
	wd_log.SetLogLineDeep(wd_log.DefaultExtLogLineMaxDeep)
	pkgJson.InitPkgJsonContent(woodpecker_plugin_env.PackageJson)

	// register helpers once
	//wd_template.RegisterSettings(wd_template.DefaultHelpers)

	app := cli.NewCliApp()

	// kubernetes runner patch
	if _, err := os.Stat("/run/drone/env"); err == nil {
		errDotEnv := godotenv.Overload("/run/drone/env")
		if errDotEnv != nil {
			wd_log.Fatalf("load /run/drone/env err: %v", errDotEnv)
		}
	}

	args := os.Args
	if err := app.Run(args); nil != err {
		color.Redf("cli err at %v\n", err)
	}
}
