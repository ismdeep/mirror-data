#!/usr/bin/env bash

set -e

target="${1:?}"

apt-get update
apt-get upgrade -y

cd /src/
make "${target}"
