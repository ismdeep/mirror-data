#!/usr/bin/env bash

set -e

docker run --rm -v "$PWD:/src/" -it python:3.10-bookworm bash -c 'set -e;cd /src/;make all'
