package plugin_test

import (
	"encoding/json"
	"github.com/woodpecker-kit/woodpecker-plugin-env/plugin"
	"github.com/woodpecker-kit/woodpecker-tools/wd_info"
	"github.com/woodpecker-kit/woodpecker-tools/wd_log"
	"github.com/woodpecker-kit/woodpecker-tools/wd_mock"
	"github.com/woodpecker-kit/woodpecker-tools/wd_short_info"
	"github.com/woodpecker-kit/woodpecker-tools/wd_steps_transfer"
	"testing"
)

func TestCheckArgsPlugin(t *testing.T) {
	t.Log("mock Plugin")
	p := mockPluginWithStatus(t, wd_info.BuildStatusSuccess)

	// statusSuccess
	var statusSuccess plugin.Plugin
	deepCopyByPlugin(&p, &statusSuccess)

	// statusNotSupport
	var statusNotSupport plugin.Plugin
	deepCopyByPlugin(&p, &statusNotSupport)
	statusNotSupport.WoodpeckerInfo = wd_mock.NewWoodpeckerInfo(
		wd_mock.WithCurrentPipelineStatus("not_support"),
	)

	tests := []struct {
		name              string
		p                 plugin.Plugin
		isDryRun          bool
		workRoot          string
		wantArgFlagNotErr bool
	}{
		{
			name:              "statusSuccess",
			p:                 statusSuccess,
			wantArgFlagNotErr: true,
		},
		{
			name: "statusNotSupport",
			p:    statusNotSupport,
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			errPluginRun := tc.p.Exec()
			if tc.wantArgFlagNotErr {
				if errPluginRun != nil {
					wdShotInfo := wd_short_info.ParseWoodpeckerInfo2Short(*tc.p.WoodpeckerInfo)
					wd_log.VerboseJsonf(wdShotInfo, "print WoodpeckerInfoShort")
					wd_log.VerboseJsonf(tc.p.Config, "print Config")
					t.Fatalf("wantArgFlagNotErr %v\np.Exec() error:\n%v", tc.wantArgFlagNotErr, errPluginRun)
					return
				}
			} else {
				if errPluginRun == nil {
					t.Fatalf("test case [ %s ], wantArgFlagNotErr %v, but p.Exec() not error", tc.name, tc.wantArgFlagNotErr)
				}
				t.Logf("check args error: %v", errPluginRun)
			}
		})
	}
}

func TestPlugin(t *testing.T) {
	t.Log("do Plugin")
	if envCheck(t) {
		return
	}
	if envMustArgsCheck(t) {
		return
	}
	t.Log("mock Plugin")
	p := mockPluginWithStatus(t, wd_info.BuildStatusSuccess)
	//wd_log.VerboseJsonf(p, "print plugin info")

	t.Log("mock plugin config")

	// remove or change this code
	p.Config.PaddingLeftMax = envPaddingLeftMax
	p.Config.EnvPrintKeys = envPrinterPrintKeys

	// statusSuccess
	var statusSuccess plugin.Plugin
	deepCopyByPlugin(&p, &statusSuccess)

	// statusFailure
	var statusFailure plugin.Plugin
	deepCopyByPlugin(&p, &statusFailure)
	statusFailure.WoodpeckerInfo = wd_mock.NewWoodpeckerInfo(
		wd_mock.WithCurrentPipelineStatus(wd_info.BuildStatusFailure),
	)

	// tagPipeline
	var tagPipeline plugin.Plugin
	deepCopyByPlugin(&p, &tagPipeline)
	tagPipeline.WoodpeckerInfo = wd_mock.NewWoodpeckerInfo(
		wd_mock.WithFastMockTag("v1.0.0", "new tag"),
	)

	// pullRequestPipeline
	var pullRequestPipeline plugin.Plugin
	deepCopyByPlugin(&p, &pullRequestPipeline)
	pullRequestPipeline.WoodpeckerInfo = wd_mock.NewWoodpeckerInfo(
		wd_mock.WithFastMockPullRequest("1", "new pr", "feature-support", "main", "main"),
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
			name: "statusSuccess",
			p:    statusSuccess,
		},
		{
			name:     "statusFailure",
			p:        statusFailure,
			isDryRun: true,
		},
		{
			name:     "tagPipeline",
			p:        tagPipeline,
			isDryRun: true,
		},
		{
			name:     "pullRequestPipeline",
			p:        pullRequestPipeline,
			isDryRun: true,
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			tc.p.Config.DryRun = tc.isDryRun
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

func mockPluginWithStatus(t *testing.T, status string) plugin.Plugin {
	p := plugin.Plugin{
		Name:    mockName,
		Version: mockVersion,
	}
	// use env:PLUGIN_DEBUG
	p.Config.Debug = valEnvPluginDebug
	p.Config.TimeoutSecond = envTimeoutSecond
	p.Config.RootPath = testGoldenKit.GetTestDataFolderFullPath()
	p.Config.StepsTransferPath = wd_steps_transfer.DefaultKitStepsFileName

	// mock woodpecker info
	//t.Log("mockPluginWithStatus")
	woodpeckerInfo := wd_mock.NewWoodpeckerInfo(
		wd_mock.WithCurrentPipelineStatus(status),
	)
	p.WoodpeckerInfo = woodpeckerInfo

	// mock all config at here

	return p
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
