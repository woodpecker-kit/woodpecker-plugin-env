package plugin

const ()

type (
	// Config plugin private config
	Config struct {
		Debug             bool
		TimeoutSecond     uint
		RootPath          string
		StepsTransferPath string

		EnvPrintKeys      []string
		PaddingLeftMax    int
		StepsTransferDemo bool
	}
)
