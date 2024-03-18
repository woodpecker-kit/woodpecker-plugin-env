[![ci](https://github.com/woodpecker-kit/woodpecker-plugin-env/workflows/ci/badge.svg)](https://github.com/woodpecker-kit/woodpecker-plugin-env/actions/workflows/ci.yml)

[![go mod version](https://img.shields.io/github/go-mod/go-version/woodpecker-kit/woodpecker-plugin-env?label=go.mod)](https://github.com/woodpecker-kit/woodpecker-plugin-env)
[![GoDoc](https://godoc.org/github.com/woodpecker-kit/woodpecker-plugin-env?status.png)](https://godoc.org/github.com/woodpecker-kit/woodpecker-plugin-env)
[![goreportcard](https://goreportcard.com/badge/github.com/woodpecker-kit/woodpecker-plugin-env)](https://goreportcard.com/report/github.com/woodpecker-kit/woodpecker-plugin-env)

[![GitHub license](https://img.shields.io/github/license/woodpecker-kit/woodpecker-plugin-env)](https://github.com/woodpecker-kit/woodpecker-plugin-env)
[![codecov](https://codecov.io/gh/woodpecker-kit/woodpecker-plugin-env/branch/main/graph/badge.svg)](https://codecov.io/gh/woodpecker-kit/woodpecker-plugin-env)
[![GitHub latest SemVer tag)](https://img.shields.io/github/v/tag/woodpecker-kit/woodpecker-plugin-env)](https://github.com/woodpecker-kit/woodpecker-plugin-env/tags)
[![GitHub release)](https://img.shields.io/github/v/release/woodpecker-kit/woodpecker-plugin-env)](https://github.com/woodpecker-kit/woodpecker-plugin-env/releases)

## for what

- this project used to woodpecker plugin

## Contributing

[![Contributor Covenant](https://img.shields.io/badge/contributor%20covenant-v1.4-ff69b4.svg)](.github/CONTRIBUTING_DOC/CODE_OF_CONDUCT.md)
[![GitHub contributors](https://img.shields.io/github/contributors/woodpecker-kit/woodpecker-plugin-env)](https://github.com/woodpecker-kit/woodpecker-plugin-env/graphs/contributors)

We welcome community contributions to this project.

Please read [Contributor Guide](.github/CONTRIBUTING_DOC/CONTRIBUTING.md) for more information on how to get started.

请阅读有关 [贡献者指南](.github/CONTRIBUTING_DOC/zh-CN/CONTRIBUTING.md) 以获取更多如何入门的信息

## Features

- [ ] more perfect test case coverage
- [ ] more perfect benchmark case

## usage

- use this template, replace list below and add usage
    - `github.com/woodpecker-kit/woodpecker-plugin-env` to your package name
    - `woodpecker-kit` to your owner name
    - `woodpecker-plugin-env` to your project name

- use github action for this workflow push to docker hub, must add at github secrets 
    - `DOCKERHUB_OWNER` user of docker hub
    - `DOCKERHUB_REPO_NAME` repo name of docker hub
    - `DOCKERHUB_TOKEN` token of docker hub user

- if use `wd_steps_transfer` just add `.woodpecker_kit.steps.transfer` at git ignore

### workflow usage

- workflow with backend `docker`

```yml
labels:
  backend: docker
steps:
  woodpecker-plugin-env:
    image: sinlov/woodpecker-plugin-env:latest
    pull: false
    settings:
      # debug: true
      # not-empty-envs: # check env not empty v1.7.+ support
        # - WOODPECKER_AGENT_USER_HOME
      env-printer-print-keys: # print env keys
        - GOPATH
        - GOPRIVATE
        - GOBIN
      # env-printer-padding-left-max: # padding left max
        ## https://woodpecker-ci.org/docs/usage/secrets
        # from_secret: secret_printer_padding_left_max
      steps-transfer-demo: false # open this show steps transfer demo
```

- workflow with backend `local`, must install at local and effective at evn `PATH`
- install at ${GOPATH}/bin, latest

```bash
go install -a github.com/woodpecker-kit/woodpecker-plugin-env/cmd/woodpecker-plugin-env@latest
```

- install at ${GOPATH}/bin, v1.0.0

```bash
go install -v github.com/woodpecker-kit/woodpecker-plugin-env/cmd/woodpecker-plugin-env@v1.0.0
```

```yml
labels:
  backend: local
steps:
  woodpecker-plugin-env:
    image: woodpecker-plugin-env
    settings:
      # debug: false
      # not-empty-envs: # check env not empty v1.7.+ support
      # - WOODPECKER_AGENT_USER_HOME
      env-printer-print-keys: # print env keys
        - GOPATH
        - GOPRIVATE
        - GOBIN
      env-printer-padding-left-max: 36 # padding left max
      steps-transfer-demo: false # open this show steps transfer demo
```

### settings.debug

- if open `settings.debug` will try file browser use `override` for debug.
- if open `settings.woodpecker_kit_steps_transfer_disable_out` will disable out of `wd_steps_transfer`

---

- want dev this project, see [doc](doc/README.md)