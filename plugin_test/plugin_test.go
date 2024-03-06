package plugin_test

import (
	"encoding/json"
	"github.com/woodpecker-kit/woodpecker-plugin-env/plugin"
	"github.com/woodpecker-kit/woodpecker-tools/wd_info"
	"github.com/woodpecker-kit/woodpecker-tools/wd_log"
	"github.com/woodpecker-kit/woodpecker-tools/wd_mock"
	"github.com/woodpecker-kit/woodpecker-tools/wd_steps_transfer"
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
	p.Config.RootPath = testGoldenKit.GetTestDataFolderFullPath()
	p.Config.StepsTransferPath = wd_steps_transfer.DefaultKitStepsFileName

	// remove or change this code
	p.Config.PaddingLeftMax = envPaddingLeftMax
	p.Config.EnvPrintKeys = envPrinterPrintKeys

	// mock woodpecker info
	woodpeckerInfo := wd_mock.NewWoodpeckerInfo(
		wd_mock.WithCurrentPipelineStatus(wd_info.BuildStatusCreated),
	)
	p.WoodpeckerInfo = woodpeckerInfo

	// statusSuccess
	var statusSuccess plugin.Plugin
	deepCopyByPlugin(&p, &statusSuccess)

	// statusFailure
	var statusFailure plugin.Plugin
	deepCopyByPlugin(&p, &statusFailure)
	statusFailure.WoodpeckerInfo = wd_mock.NewWoodpeckerInfo(
		wd_mock.WithCurrentPipelineStatus(wd_info.BuildStatusFailure),
	)

	tests := []struct {
		name            string
		p               plugin.Plugin
		isDryRun        bool
		workRoot        string
		ossTransferKey  string
		ossTransferData interface{}
		wantErr         bool
	}{
		{
			name:     "statusSuccess",
			p:        statusSuccess,
			isDryRun: true,
		},
		{
			name:     "statusFailure",
			p:        statusFailure,
			isDryRun: true,
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if tc.workRoot != "" {
				tc.p.Config.RootPath = tc.workRoot
				errGenTransferData := generateTransferStepsOut(
					tc.p,
					tc.ossTransferKey,
					tc.ossTransferData,
				)
				if errGenTransferData != nil {
					t.Fatal(errGenTransferData)
				}
			}
			err := tc.p.Exec()
			if (err != nil) != tc.wantErr {
				t.Errorf("FeishuPlugin.Exec() error = %v, wantErr %v", err, tc.wantErr)
				return
			}
		})
	}
}

func deepCopyByPlugin(src, dst *plugin.Plugin) {
	if tmp, err := json.Marshal(&src); err != nil {
		return
	} else {
		err = json.Unmarshal(tmp, dst)
		return
	}
}

func generateTransferStepsOut(plugin plugin.Plugin, mark string, data interface{}) error {
	_, err := wd_steps_transfer.Out(plugin.Config.RootPath, plugin.Config.StepsTransferPath, *plugin.WoodpeckerInfo, mark, data)
	return err
}
