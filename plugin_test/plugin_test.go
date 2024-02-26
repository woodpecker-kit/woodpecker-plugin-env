package plugin_test

import (
	"github.com/woodpecker-kit/woodpecker-plugin-env/plugin"
	"github.com/woodpecker-kit/woodpecker-tools/wd_info"
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

	t.Log("mock woodpecker info")

	// use env:ENV_DEBUG
	p.Config.Debug = envDebug
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
