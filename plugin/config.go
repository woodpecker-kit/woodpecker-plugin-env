package plugin

const ()

type (
	// Config plugin private config
	Config struct {
		Debug         bool
		TimeoutSecond uint

		EnvPrintKeys   []string
		PaddingLeftMax int
	}
)
