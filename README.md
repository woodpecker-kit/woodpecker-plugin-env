[![docker hub version semver](https://img.shields.io/docker/v/sinlov/woodpecker-plugin-env?sort=semver)](https://hub.docker.com/r/sinlov/woodpecker-plugin-env/tags?page=1&ordering=last_updated)
[![docker hub image size](https://img.shields.io/docker/image-size/sinlov/woodpecker-plugin-env)](https://hub.docker.com/r/sinlov/woodpecker-plugin-env)
[![docker hub image pulls](https://img.shields.io/docker/pulls/sinlov/woodpecker-plugin-env)](https://hub.docker.com/r/sinlov/woodpecker-plugin-env/tags?page=1&ordering=last_updated)

[![ci](https://github.com/woodpecker-kit/woodpecker-plugin-env/workflows/ci/badge.svg?)](https://github.com/woodpecker-kit/woodpecker-plugin-env/actions/workflows/ci.yml)

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

### workflow usage

- workflow with backend `docker`

```yml
labels:
  backend: docker
steps:
  env:
    image: sinlov/woodpecker-plugin-env:latest
    pull: false
    settings:
      # debug: true
      env_printer_print_keys: # print env keys
        - GOPATH
        - GOPRIVATE
        - GOBIN
      env_printer_padding_left_max: 36 # padding left max
```

- workflow with backend `local`, must install at local and effective at evn `PATH`

```bash
# install at ${GOPATH}/bin
$ go install -v github.com/woodpecker-kit/woodpecker-plugin-env/cmd/woodpecker-plugin-env@latest
# install version v1.0.0
$ go install -v github.com/woodpecker-kit/woodpecker-plugin-env/cmd/woodpecker-plugin-env@v1.0.0
```

```yml
labels:
  backend: local
steps:
  env:
    image: woodpecker-plugin-env
    settings:
      # debug: false
      env_printer_print_keys: # print env keys
        - GOPATH
        - GOPRIVATE
        - GOBIN
      env_printer_padding_left_max: 36 # padding left max
```

---

- want dev this project, see [doc](doc/README.md)