## usage

- use this template, replace list below and add usage
    - `github.com/woodpecker-kit/woodpecker-plugin-env` to your package name
    - `woodpecker-kit` to your owner name
    - `woodpecker-plugin-env` to your project name

- use github action for this workflow push to docker hub, must add
    - variables `ENV_DOCKERHUB_OWNER` user of docker hub
    - variables `ENV_DOCKERHUB_REPO_NAME` repo name of docker hub
    - secrets `DOCKERHUB_TOKEN` token of docker hub user from [hub.docker](https://hub.docker.com/settings/security)

- check `docker-bake.hcl` config, change to your docker image

- if you use `wd_steps_transfer` just add `.woodpecker_kit.steps.transfer` at git ignore
- change code start with `// change or remove`

# dev

## just start dev

- minimum go version: go 1.21
- change `go 1.21`, `^1.21`, `1.21.13` to new go version
- lint
    - golangci-lint will update to v2
    - change `golangci/golangci-lint:v2.1.6` # https://hub.docker.com/r/golangci/golangci-lint/tags
    - change `woodpeckerci/plugin-reviewdog-golangci-lint:1.61.0` version
      from [woodpeckerci/plugin-reviewdog-golangci-lint](https://hub.docker.com/r/woodpeckerci/plugin-reviewdog-golangci-lint/tags)
      code [woodpecker-plugins/reviewdog-golangci-lint](https://codeberg.org/woodpecker-plugins/reviewdog-golangci-lint)

- [reviewdog](https://github.com/reviewdog/reviewdog) code review tool integrated with any code analysis tools

- change goreleaser
    - from [goreleaser docker tags](https://hub.docker.com/r/goreleaser/goreleaser/tags) to new version
    - `golangci/golangci-lint-action@v8`
    - from [goreleaser version release](https://github.com/goreleaser/goreleaser/releases) to new version


in go mod project

```bash
# warning use private git host must set
# global set for once
# add private git host like github.com to evn GOPRIVATE
$ go env -w GOPRIVATE='github.com'
# use ssh proxy
# set ssh-key to use ssh as http
$ git config --global url."git@github.com:".insteadOf "https://github.com/"
# or use PRIVATE-TOKEN
# set PRIVATE-TOKEN as gitlab or gitea
$ git config --global http.extraheader "PRIVATE-TOKEN: {PRIVATE-TOKEN}"
# set this rep to download ssh as https use PRIVATE-TOKEN
$ git config --global url."ssh://github.com/".insteadOf "https://github.com/"

# before above global settings
# test version info
$ git ls-remote -q https://github.com/woodpecker-kit/woodpecker-plugin-env.git

# test depends see full version
$ go list -mod readonly -v -m -versions github.com/woodpecker-kit/woodpecker-plugin-env
# or use last version add go.mod by script
$ echo "go mod edit -require=$(go list -mod=readonly -m -versions github.com/woodpecker-kit/woodpecker-plugin-env | awk '{print $1 "@" $NF}')"
$ echo "go mod vendor"
```

### libs

| lib                                       | version |
|:------------------------------------------|:--------|
| https://github.com/stretchr/testify       | v1.10.0 |
| https://github.com/sebdah/goldie          | v2.5.5  |
| https://github.com/sinlov-go/unittest-kit | v1.2.1  |

- more libs see main
  branch [go.mod](https://github.com/woodpecker-kit/woodpecker-plugin-env/blob/main/go.mod)

## dev tasks

```bash
# It needs to be executed after the first use or update of dependencies.
make init dep
```

- test code

```bash
make test
# benchmark and coverage show
make ci.test.benchmark ci.coverage.show
```

- ci to fast check as CI pipeline

```bash
# check style at local
make style

# run ci at local
make ci
```

### docker

```bash
# then test build as test/Dockerfile
$ make dockerTestRestartLatest
# clean test build
$ make dockerTestPruneLatest

# more info see
$ make helpDocker
```
