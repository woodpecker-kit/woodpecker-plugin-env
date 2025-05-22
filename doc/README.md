# woodpecker-plugin-env

### EngineeringStructure

```
.
├── Dockerfile                     # ci docker build
├── build.dockerfile               # local docker build
├── Makefile                       # make entry
├── README.md
├── build                          # build output folder
├── dist                           # dist output folder
├── cmd
│     ├── cli
│     │     ├── app.go             # cli entry
│     │     ├── cli_aciton_test.go # cli action test
│     │     └── cli_action.go      # cli action
│     └── woodpecker-plugin-env    # command line main package install and dev entrance
│         ├── main.go                   # command line entry
│         └── main_test.go              # integrated test entry
├── constant                       # constant package
│         ├── common_flag.go         # common environment variable
│         ├── platform_flag.go       # platform environment variable
│         └── version.go             # semver version constraint set
├── doc
│         ├── README.md              # command line tools documentation
│         └── docs.md                # woodpecker documentation
├── go.mod
├── go.sum
├── package.json                   # command line profile information for embed
├── resource.go                    # embed resource
├── internal                          # toolkit package
│         ├── pkgJson                 # package.json toolkit
│         └── version_check           # version check by semver
├── plugin                         # plugin package
│         ├── flag.go                 # plugin flag
│         ├── impl.go                 # plugin implement
│         ├── plugin.go               # plugin entry
│         └── settings.go             # plugin settings
├── plugin_test                    # plugin test
│         ├── init_test.go            # each test init
│         └── plugin_test.go          # plugin test
├── z-MakefileUtils                # make toolkit
└── zymosis                         # resource mark by https://github.com/convention-change/zymosis


```

### log

- open debug log by env `PLUGIN_DEBUG=true` or global flag `--plugin.debug true`

```go
package foo

import (
	"github.com/sinlov-go/unittest-kit/env_kit"
	"github.com/urfave/cli/v2"
	"github.com/woodpecker-kit/woodpecker-tools/wd_log"
	"github.com/woodpecker-kit/woodpecker-tools/wd_urfave_cli_v2"
	"os/user"
)

func GlobalBeforeAction(c *cli.Context) error {
	isDebug := wd_urfave_cli_v2.IsBuildDebugOpen(c)
	if isDebug {
		// print global debug info
		allEnvPrintStr := env_kit.FindAllEnv4PrintAsSortJust(36)
		wd_log.Verbosef("==> plugin start with all env:\n%s", allEnvPrintStr)
		currentUser, err := user.Current()
		if err == nil {
			wd_log.Verbosef("==> current Username : %s\n", currentUser.Username)
			wd_log.Verbosef("==> current user name: %s\n", currentUser.Name)
			wd_log.Verbosef("==> current gid: %s, uid: %s\n", currentUser.Gid, currentUser.Uid)
			wd_log.Verbosef("==> current user home: %s\n", currentUser.HomeDir)
		}
		wd_log.OpenDebug()
	}
	return nil
}
```

### template

- [https://github.com/aymerick/raymond](https://github.com/aymerick/raymond)
- function doc [https://masterminds.github.io/sprig/](https://masterminds.github.io/sprig/)

- open template support at cli `main.go`

```go
package main

func main() {
	// register helpers once
	wd_template.RegisterSettings(wd_template.DefaultHelpers)
}
```

- and open at test `init_test.go`

```go
package plugin_test

func init() {
	// if open wd_template please open this
	wd_template.RegisterSettings(wd_template.DefaultHelpers)
}
```