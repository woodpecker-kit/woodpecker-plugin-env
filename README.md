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

- see [doc](doc/docs.md)

## Notice

- want dev this project, see [doc](doc/README.md)