//go:build !test

package main

import (
	os "os"

	"github.com/gookit/color"
	"github.com/joho/godotenv"
	woodpecker_plugin_env "github.com/woodpecker-kit/woodpecker-plugin-env"
	"github.com/woodpecker-kit/woodpecker-plugin-env/cmd/cli"
	"github.com/woodpecker-kit/woodpecker-plugin-env/constant"
	"github.com/woodpecker-kit/woodpecker-plugin-env/internal/cli_kit/pkg_kit"
	"github.com/woodpecker-kit/woodpecker-tools/wd_log"
)

//nolint:gochecknoglobals
var (
	// Populated by goreleaser during build.
	version    = "unknown"
	rawVersion = "unknown"
	buildID    string
	commit     = "?"
	date       = ""
)

func init() {
	if buildID == "" {
		buildID = "unknown"
	}
}

func main() {
	wd_log.SetLogLineDeep(wd_log.DefaultExtLogLineMaxDeep)
	pkg_kit.InitPkgJsonContent(woodpecker_plugin_env.PackageJson)

	bdInfo := pkg_kit.NewBuildInfo(
		pkg_kit.GetPackageJsonName(),
		pkg_kit.GetPackageJsonDescription(),
		version,
		rawVersion,
		buildID,
		commit,
		date,
		pkg_kit.GetPackageJsonAuthor().Name,
		constant.CopyrightStartYear,
	)

	// register helpers once
	// wd_template.RegisterSettings(wd_template.DefaultHelpers)

	// kubernetes runner patch
	if _, err := os.Stat("/run/drone/env"); err == nil {
		errDotEnv := godotenv.Overload("/run/drone/env")
		if errDotEnv != nil {
			wd_log.Fatalf("load /run/drone/env err: %v", errDotEnv)
		}
	}

	// load env file by env `PLUGIN_ENV_FILE`
	if envFile, set := os.LookupEnv("PLUGIN_ENV_FILE"); set {
		errLoadEnvFile := godotenv.Overload(envFile)
		if errLoadEnvFile != nil {
			wd_log.Fatalf("load env file %s err: %v", envFile, errLoadEnvFile)
		}
	}

	app := cli.NewCliApp(bdInfo)

	args := os.Args
	if err := app.Run(args); nil != err {
		color.Redf("cli err at %v\n", err)
	}
}
