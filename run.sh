#!/usr/bin/env bash

set -e

target="${1:?}"

git pull

docker run --rm -v "$PWD:/src/" golang:1.20-bookworm bash /src/run-in-docker.sh "${target}"

git add .
git commit -am "docs: update"
git push
