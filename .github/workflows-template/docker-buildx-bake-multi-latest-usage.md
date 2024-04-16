## need `New repository secret`

- file `docker-image-latest.yml`
- `DOCKERHUB_TOKEN` from [hub.docker](https://hub.docker.com/settings/security)
    - if close push remote can pass `DOCKERHUB_TOKEN` setting
- `DOCKERHUB_OWNER` for docker hub user
- `DOCKERHUB_REPO_NAME` for docker hub repo name

## usage at github action

```yml
  docker-buildx-bake-multi-latest:
    name: docker-buildx-bake-multi-latest
    needs:
      - version
    uses: ./.github/workflows/docker-buildx-bake-multi-latest.yml
    # if: ${{ github.event.pull_request.merged == true }}
    secrets:
      DOCKERHUB_OWNER: "${{ secrets.DOCKERHUB_OWNER }}"
      DOCKERHUB_REPO_NAME: "${{ secrets.DOCKERHUB_REPO_NAME }}"
      DOCKERHUB_TOKEN: "${{ secrets.DOCKERHUB_TOKEN }}"
    with:
      ghcr_package_owner_name: ${{ github.repository_owner }} # required for ghcr.io
      # push_remote_flag: ${{ github.event.pull_request.merged == true }}
      push_remote_flag: true

```