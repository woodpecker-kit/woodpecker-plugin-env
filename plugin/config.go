package plugin

const ()

type (
	// Config plugin private config
	Config struct {
		Debug             bool
		TimeoutSecond     uint
		StepsTransferPath string
		RootPath          string

		// remove or change this config
		EnvPrintKeys      []string
		PaddingLeftMax    int
		StepsTransferDemo bool
	}
)
