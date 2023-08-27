#!/usr/bin/env bash

set -e

target="${1:?}"

docker build -t ismdeep/mirror-data-runner:latest .

git pull

docker run --rm -v "$PWD:/src/" ismdeep/mirror-data-runner:latest bash /src/run-in-docker.sh "${target}"

git add .
git commit -am "docs: update" || true
git push
