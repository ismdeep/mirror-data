#!/usr/bin/env bash

set -e

# prepare version info
version="$(cat VERSION)"
go_version=$(go version)
commit_id=$(git log -1 --pretty=format:"%H" | awk '{print $1}')
commit_date=$(git log -1 --pretty=format:"%ci")


# build
build() {
  os="${1:?}"
  arch="${2:?}"
  ldflags="\
  -s -w \
  -X 'github.com/ismdeep/mirror-data/internal/version.Version=${version}' \
  -X 'github.com/ismdeep/mirror-data/internal/version.CommitDate=${commit_date}' \
  -X 'github.com/ismdeep/mirror-data/internal/version.CommitID=${commit_id}' \
  -X 'github.com/ismdeep/mirror-data/internal/version.GoVersion=${go_version}' \
  -X 'github.com/ismdeep/mirror-data/internal/version.OS=${os}' \
  -X 'github.com/ismdeep/mirror-data/internal/version.Arch=${arch}' \
  "
  GOOS=${os}  GOARCH=${arch} go build -ldflags "${ldflags}" -o "./bin/mirror-data-${version}-${os}-${arch}"  github.com/ismdeep/mirror-data
  echo "==> ./bin/mirror-data-${version}-${os}-${arch}"
}

#### MAIN ####
mkdir -p  ./bin/
rm    -rf ./bin/
mkdir -p  ./bin/

build linux  amd64
build linux  arm64
build darwin amd64
build darwin arm64