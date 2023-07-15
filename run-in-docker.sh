#!/usr/bin/env bash

set -e

# touch rclone.conf
mkdir -p /root/.config/rclone/
touch    /root/.config/rclone/rclone.conf

# install rclone
cd "$(mktemp -d)"
curl -LO https://downloads.rclone.org/v1.63.0/rclone-v1.63.0-linux-amd64.zip
unzip rclone-v1.63.0-linux-amd64.zip
cp -v rclone-v1.63.0-linux-amd64/rclone /usr/bin/rclone
chmod +x /usr/bin/rclone

# process
cd /src/
make all