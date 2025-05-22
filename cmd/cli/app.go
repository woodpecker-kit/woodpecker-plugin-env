package cli

import (
	"github.com/urfave/cli/v2"
	"github.com/woodpecker-kit/woodpecker-plugin-env/constant"
	"github.com/woodpecker-kit/woodpecker-plugin-env/internal/cli_kit/pkg_kit"
	"github.com/woodpecker-kit/woodpecker-plugin-env/plugin"
	"github.com/woodpecker-kit/woodpecker-tools/wd_urfave_cli_v2"
	"github.com/woodpecker-kit/woodpecker-tools/wd_urfave_cli_v2/cli_exit_urfave"
)

const (
	defaultExitCode = 1
)

func NewCliApp(bdInfo pkg_kit.BuildInfo) *cli.App {
	cli_exit_urfave.ChangeDefaultExitCode(defaultExitCode)

	app := cli.NewApp()
	app.EnableBashCompletion = true
	app.Name = bdInfo.PgkNameString()
	app.Version = pkg_kit.FetchNowVersion()

	if pkg_kit.GetPackageJsonHomepage() != "" {
		app.Usage = "see: " + pkg_kit.GetPackageJsonHomepage()
	}

	app.Description = pkg_kit.GetPackageJsonDescription()

	jsonAuthor := pkg_kit.GetPackageJsonAuthor()
	author := &cli.Author{
		Name:  jsonAuthor.Name,
		Email: jsonAuthor.Email,
	}
	app.Authors = []*cli.Author{
		author,
	}
	app.Copyright = bdInfo.String()

	flags := wd_urfave_cli_v2.UrfaveCliAppendCliFlags(
		wd_urfave_cli_v2.WoodpeckerUrfaveCliFlags(),
		constant.CommonFlag(),
		constant.HideCommonGlobalFlag(),
		plugin.GlobalFlag(),
		plugin.HideGlobalFlag(),
	)

	app.Flags = flags
	app.Before = GlobalBeforeAction
	app.Action = GlobalAction
	app.After = GlobalAfterAction

	return app
}
