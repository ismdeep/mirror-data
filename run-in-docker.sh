#!/usr/bin/env bash

set -e

target="${1}"

if [ "${target}" == "" ]; then
  target="all"
fi

docker run --rm -v "$PWD:/src/" -it python:3.10-bookworm bash -c 'set -e;cd /src/;make '"${target}"''
