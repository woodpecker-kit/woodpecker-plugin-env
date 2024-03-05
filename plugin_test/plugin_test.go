package plugin_test

import (
	"github.com/woodpecker-kit/woodpecker-plugin-env/plugin"
	"github.com/woodpecker-kit/woodpecker-tools/wd_info"
	"github.com/woodpecker-kit/woodpecker-tools/wd_log"
	"github.com/woodpecker-kit/woodpecker-tools/wd_mock"
	"testing"
)

func TestPlugin(t *testing.T) {
	t.Log("mock Plugin")

	p := plugin.Plugin{
		Name:    mockName,
		Version: mockVersion,
	}

	t.Log("do Plugin")
	if envCheck(t) {
		return
	}
	wd_log.VerboseJsonf(p, "print plugin info")
	if envMustArgsCheck(t) {
		return
	}

	t.Log("mock woodpecker info")

	// use env:PLUGIN_DEBUG
	p.Config.Debug = valEnvPluginDebug
	p.Config.TimeoutSecond = envTimeoutSecond

	// remove or change this code
	p.Config.PaddingLeftMax = envPaddingLeftMax
	p.Config.EnvPrintKeys = envPrinterPrintKeys

	// mock woodpecker info
	woodpeckerInfo := wd_mock.NewWoodpeckerInfo(
		wd_mock.WithCurrentPipelineStatus(wd_info.BuildStatusCreated),
	)
	p.WoodpeckerInfo = woodpeckerInfo

	err := p.Exec()

	t.Log("verify woodpecker Plugin")
	if err != nil {
		t.Fatal(err)
	}
}
