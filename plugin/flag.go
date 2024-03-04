package plugin

import (
	"github.com/urfave/cli/v2"
	"github.com/woodpecker-kit/woodpecker-tools/wd_flag"
	"github.com/woodpecker-kit/woodpecker-tools/wd_info"
	"github.com/woodpecker-kit/woodpecker-tools/wd_log"
)

const (
	// remove or change this code

	CliNamePrinterPrintKeys = "settings.env_printer_print_keys"
	EnvPrinterPrintKeys     = "PLUGIN_ENV_PRINTER_PRINT_KEYS"

	CliNamePrinterPaddingLeftMax = "settings.env_printer_padding_left_max"
	EnvPrinterPaddingLeftMax     = "PLUGIN_ENV_PRINTER_PADDING_LEFT_MAX"

	CliNameStepsTransferDemo = "settings.steps_transfer_demo"
	EnvStepsTransferDemo     = "PLUGIN_STEPS_TRANSFER_DEMO"
)

// GlobalFlag
// Other modules also have flags
func GlobalFlag() []cli.Flag {
	return []cli.Flag{
		// new flag string template if no use, please replace this
		&cli.StringSliceFlag{
			Name:    CliNamePrinterPrintKeys,
			Usage:   "if use this args, will print env by keys",
			EnvVars: []string{EnvPrinterPrintKeys},
		},
		&cli.IntFlag{
			Name:    CliNamePrinterPaddingLeftMax,
			Usage:   "set env printer padding left max count, minimum 24, default 32",
			EnvVars: []string{EnvPrinterPaddingLeftMax},
			Value:   32,
		},
		&cli.BoolFlag{
			Name:    CliNameStepsTransferDemo,
			Usage:   "if use this args, will print steps transfer demo",
			EnvVars: []string{EnvStepsTransferDemo},
		},
		// env_printer_plugin end
		//&cli.StringFlag{
		//	Name:    "settings.new_arg",
		//	Usage:   "",
		//	EnvVars: []string{"PLUGIN_new_arg"},
		//},
	}
}

func HideGlobalFlag() []cli.Flag {
	return []cli.Flag{}
}

func BindCliFlags(c *cli.Context, cliName, cliVersion string, wdInfo *wd_info.WoodpeckerInfo, rootPath, stepsTransferPath string) (*Plugin, error) {
	debug := isBuildDebugOpen(c)

	config := Config{
		Debug:             debug,
		TimeoutSecond:     c.Uint(wd_flag.NameCliPluginTimeoutSecond),
		StepsTransferPath: stepsTransferPath,
		RootPath:          rootPath,

		// remove or change this code
		EnvPrintKeys:      c.StringSlice(CliNamePrinterPrintKeys),
		PaddingLeftMax:    c.Int(CliNamePrinterPaddingLeftMax),
		StepsTransferDemo: c.Bool(CliNameStepsTransferDemo),
	}

	// set default TimeoutSecond
	if config.TimeoutSecond == 0 {
		config.TimeoutSecond = 10
	}
	// set default PaddingLeftMax
	if config.PaddingLeftMax < 24 {
		config.PaddingLeftMax = 24
	}

	wd_log.Debugf("args %s: %v", wd_flag.NameCliPluginTimeoutSecond, config.TimeoutSecond)

	p := Plugin{
		Name:           cliName,
		Version:        cliVersion,
		WoodpeckerInfo: wdInfo,
		Config:         config,
	}

	return &p, nil
}

// isBuildDebugOpen
// when config or build open debug will open debug
func isBuildDebugOpen(c *cli.Context) bool {
	return c.Bool(wd_flag.NameCliPluginDebug)
}
