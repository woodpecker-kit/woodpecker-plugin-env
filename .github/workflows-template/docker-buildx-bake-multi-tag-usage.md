## need `New repository secret`

- file `docker-buildx-bake-multi-tag.yml`
- `DOCKERHUB_TOKEN` from [hub.docker](https://hub.docker.com/settings/security)
    - if close push remote can pass `DOCKERHUB_TOKEN` setting
- `DOCKERHUB_OWNER` for docker hub user
- `DOCKERHUB_REPO_NAME` for docker hub repo name

## usage at github action

```yml
  docker-buildx-bake-multi-tag:
    name: docker-buildx-bake-multi-tag
    needs:
      - version
    uses: ./.github/workflows/docker-buildx-bake-multi-tag.yml
    if: startsWith(github.ref, 'refs/tags/')
    secrets:
      DOCKERHUB_OWNER: "${{ secrets.DOCKERHUB_OWNER }}"
      DOCKERHUB_REPO_NAME: "${{ secrets.DOCKERHUB_REPO_NAME }}"
      DOCKERHUB_TOKEN: "${{ secrets.DOCKERHUB_TOKEN }}"
    with:
      push_remote_flag: true # set true to push, default false
      ghcr_package_owner_name: ${{ github.repository_owner }} # required for ghcr.io
```