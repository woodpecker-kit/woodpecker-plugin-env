package plugin_test

import (
	"github.com/woodpecker-kit/woodpecker-plugin-env/plugin"
	"github.com/woodpecker-kit/woodpecker-tools/wd_info"
	"github.com/woodpecker-kit/woodpecker-tools/wd_log"
	"github.com/woodpecker-kit/woodpecker-tools/wd_mock"
	"github.com/woodpecker-kit/woodpecker-tools/wd_short_info"
	"testing"
)

func TestCheckArgsPlugin(t *testing.T) {
	t.Log("mock Plugin")
	// successArgs
	successArgsWoodpeckerInfo := *wd_mock.NewWoodpeckerInfo(
		wd_mock.WithCurrentPipelineStatus(wd_info.BuildStatusSuccess),
	)
	successArgsSettings := mockPluginSettings()

	// notSupport
	notSupportWoodpeckerInfo := *wd_mock.NewWoodpeckerInfo(
		wd_mock.WithCurrentPipelineStatus("not_support"),
	)
	notSupportSettings := mockPluginSettings()

	tests := []struct {
		name           string
		woodpeckerInfo wd_info.WoodpeckerInfo
		settings       plugin.Settings
		workRoot       string

		isDryRun          bool
		wantArgFlagNotErr bool
	}{
		{
			name:              "successArgs",
			woodpeckerInfo:    successArgsWoodpeckerInfo,
			settings:          successArgsSettings,
			wantArgFlagNotErr: true,
		},
		{
			name:           "notSupport",
			woodpeckerInfo: notSupportWoodpeckerInfo,
			settings:       notSupportSettings,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			p := mockPluginWithSettings(t, tc.woodpeckerInfo, tc.settings)
			p.OnlyArgsCheck()
			errPluginRun := p.Exec()
			if tc.wantArgFlagNotErr {
				if errPluginRun != nil {
					wdShotInfo := wd_short_info.ParseWoodpeckerInfo2Short(p.GetWoodPeckerInfo())
					wd_log.VerboseJsonf(wdShotInfo, "print WoodpeckerInfoShort")
					wd_log.VerboseJsonf(p.Settings, "print Settings")
					t.Fatalf("wantArgFlagNotErr %v\np.Exec() error:\n%v", tc.wantArgFlagNotErr, errPluginRun)
					return
				}
				infoShot := p.ShortInfo()
				wd_log.VerboseJsonf(infoShot, "print WoodpeckerInfoShort")
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

	t.Log("mock plugin config")

	// statusSuccess
	statusSuccessWoodpeckerInfo := *wd_mock.NewWoodpeckerInfo(
		wd_mock.WithCurrentPipelineStatus(wd_info.BuildStatusSuccess),
	)
	statusSuccessSettings := mockPluginSettings()

	// statusFailure
	statusFailureWoodpeckerInfo := *wd_mock.NewWoodpeckerInfo(
		wd_mock.WithCurrentPipelineStatus(wd_info.BuildStatusFailure),
	)
	statusFailureSettings := mockPluginSettings()

	// tagPipeline
	tagPipelineWoodpeckerInfo := *wd_mock.NewWoodpeckerInfo(
		wd_mock.WithFastMockTag("v1.0.0", "new tag"),
	)
	tagPipelineSettings := mockPluginSettings()

	// pullRequestPipeline
	pullRequestPipelineWoodpeckerInfo := *wd_mock.NewWoodpeckerInfo(
		wd_mock.WithFastMockPullRequest("1", "new pr", "feature-support", "main", "main"),
	)
	pullRequestPipelineSettings := mockPluginSettings()

	tests := []struct {
		name           string
		woodpeckerInfo wd_info.WoodpeckerInfo
		settings       plugin.Settings
		workRoot       string

		ossTransferKey  string
		ossTransferData interface{}

		isDryRun bool
		wantErr  bool
	}{
		{
			name:           "statusSuccess",
			woodpeckerInfo: statusSuccessWoodpeckerInfo,
			settings:       statusSuccessSettings,
		},
		{
			name:           "statusFailure",
			woodpeckerInfo: statusFailureWoodpeckerInfo,
			settings:       statusFailureSettings,
			isDryRun:       true,
		},
		{
			name:           "tagPipeline",
			woodpeckerInfo: tagPipelineWoodpeckerInfo,
			settings:       tagPipelineSettings,
			isDryRun:       true,
		},
		{
			name:           "pullRequestPipeline",
			woodpeckerInfo: pullRequestPipelineWoodpeckerInfo,
			settings:       pullRequestPipelineSettings,
			isDryRun:       true,
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			p := mockPluginWithSettings(t, tc.woodpeckerInfo, tc.settings)
			p.Settings.DryRun = tc.isDryRun
			if tc.workRoot != "" {
				p.Settings.RootPath = tc.workRoot
				errGenTransferData := generateTransferStepsOut(
					p,
					tc.ossTransferKey,
					tc.ossTransferData,
				)
				if errGenTransferData != nil {
					t.Fatal(errGenTransferData)
				}
			}
			err := p.Exec()
			if (err != nil) != tc.wantErr {
				t.Errorf("FeishuPlugin.Exec() error = %v, wantErr %v", err, tc.wantErr)
				return
			}
		})
	}
}
