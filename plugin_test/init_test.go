package plugin_test

import (
	"fmt"
	"github.com/sinlov-go/unittest-kit/env_kit"
	"github.com/sinlov-go/unittest-kit/unittest_file_kit"
	"github.com/woodpecker-kit/woodpecker-plugin-env/plugin"
	"github.com/woodpecker-kit/woodpecker-tools/wd_flag"
	"github.com/woodpecker-kit/woodpecker-tools/wd_info"
	"github.com/woodpecker-kit/woodpecker-tools/wd_log"
	"github.com/woodpecker-kit/woodpecker-tools/wd_steps_transfer"
	"os"
	"path/filepath"
	"runtime"
	"testing"
)

const (
	keyEnvDebug  = "CI_DEBUG"
	keyEnvCiNum  = "CI_NUMBER"
	keyEnvCiKey  = "CI_KEY"
	keyEnvCiKeys = "CI_KEYS"

	mockVersion = "v1.0.0"
	mockName    = "woodpecker-plugin-env"
)

var (
	// testBaseFolderPath
	//  test base dir will auto get by package init()
	testBaseFolderPath = ""
	testGoldenKit      *unittest_file_kit.TestGoldenKit

	envTimeoutSecond    uint
	envPaddingLeftMax   = 0
	envPrinterPrintKeys []string

	// mustSetInCiEnvList
	//  for check set in CI env not empty
	mustSetInCiEnvList = []string{
		wd_flag.EnvKeyCiSystemPlatform,
		wd_flag.EnvKeyCiSystemVersion,
	}
	// mustSetArgsAsEnvList
	mustSetArgsAsEnvList = []string{
		//plugin.EnvStepsTransferDemo,
	}

	valEnvPluginDebug = false
)

func init() {
	testBaseFolderPath, _ = getCurrentFolderPath()
	wd_log.SetLogLineDeep(2)
	// if open wd_template please open this
	//wd_template.RegisterSettings(wd_template.DefaultHelpers)

	envTimeoutSecond = uint(env_kit.FetchOsEnvInt(wd_flag.EnvKeyPluginTimeoutSecond, 10))

	envPaddingLeftMax = env_kit.FetchOsEnvInt(plugin.EnvPrinterPaddingLeftMax, 24)
	envPrinterPrintKeys = env_kit.FetchOsEnvStringSlice(plugin.EnvPrinterPrintKeys)

	testGoldenKit = unittest_file_kit.NewTestGoldenKit(testBaseFolderPath)

	valEnvPluginDebug = env_kit.FetchOsEnvBool(wd_flag.EnvKeyPluginDebug, false)
}

// test case basic tools start
// getCurrentFolderPath
//
//	can get run path this golang dir
func getCurrentFolderPath() (string, error) {
	_, file, _, ok := runtime.Caller(1)
	if !ok {
		return "", fmt.Errorf("can not get current file info")
	}
	return filepath.Dir(file), nil
}

// test case basic tools end

func envCheck(t *testing.T) bool {

	if valEnvPluginDebug {
		wd_log.OpenDebug()
	}

	// most CI system will set env CI to true
	envCI := env_kit.FetchOsEnvStr("CI", "")
	if envCI == "" {
		t.Logf("not in CI system, skip envCheck")
		return false
	}
	t.Logf("check env for CI system")
	return env_kit.MustHasEnvSetByArray(t, mustSetInCiEnvList)
}

func envMustArgsCheck(t *testing.T) bool {
	for _, item := range mustSetArgsAsEnvList {
		if os.Getenv(item) == "" {
			t.Logf("plasee set env: %s, than run test\nfull need set env %v", item, mustSetArgsAsEnvList)
			return true
		}
	}
	return false
}

func generateTransferStepsOut(plugin plugin.Plugin, mark string, data interface{}) error {
	_, err := wd_steps_transfer.Out(plugin.Settings.RootPath, plugin.Settings.StepsTransferPath, plugin.GetWoodPeckerInfo(), mark, data)
	return err
}

func mockPluginSettings() plugin.Settings {
	// all mock settings can set here
	settings := plugin.Settings{
		// use env:PLUGIN_DEBUG
		Debug:             valEnvPluginDebug,
		TimeoutSecond:     envTimeoutSecond,
		RootPath:          testGoldenKit.GetTestDataFolderFullPath(),
		StepsTransferPath: wd_steps_transfer.DefaultKitStepsFileName,
	}

	// remove or change this code
	settings.PaddingLeftMax = envPaddingLeftMax
	settings.EnvPrintKeys = envPrinterPrintKeys

	return settings

}

func mockPluginWithSettings(t *testing.T, woodpeckerInfo wd_info.WoodpeckerInfo, settings plugin.Settings) plugin.Plugin {
	p := plugin.Plugin{
		Name:    mockName,
		Version: mockVersion,
	}

	// mock woodpecker info
	//t.Log("mockPluginWithStatus")
	p.SetWoodpeckerInfo(woodpeckerInfo)

	p.Settings = settings
	return p
}
