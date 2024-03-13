package plugin

import "github.com/woodpecker-kit/woodpecker-tools/wd_info"

const (
	// StepsTransferMarkDemoConfig
	// steps transfer key
	StepsTransferMarkDemoConfig = "demo_config"
)

var (
	// pluginBuildStateSupport
	pluginBuildStateSupport = []string{
		wd_info.BuildStatusCreated,
		wd_info.BuildStatusRunning,
		wd_info.BuildStatusSuccess,
		wd_info.BuildStatusFailure,
		wd_info.BuildStatusError,
		wd_info.BuildStatusKilled,
	}
)

type (
	// Config plugin private config
	Config struct {
		Debug             bool
		TimeoutSecond     uint
		StepsTransferPath string
		StepsOutDisable   bool
		RootPath          string

		DryRun bool

		// remove or change this config
		EnvPrintKeys      []string
		PaddingLeftMax    int
		StepsTransferDemo bool
	}
)
