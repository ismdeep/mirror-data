# This is runtime basic image
FROM golang:1.20-bookworm
RUN set -e; \
    apt-get update; \
    apt-get install -y unzip; \
    curl -LO https://github.com/rclone/rclone/releases/download/v1.63.1/rclone-v1.63.1-linux-amd64.zip; \
    unzip rclone-v1.63.1-linux-amd64.zip; \
    mv rclone-v1.63.1-linux-amd64/rclone /usr/bin/; \
    rm -rfv rclone-v1.63.1-linux-amd64; \
    rm -f   rclone-v1.63.1-linux-amd64.zip; \
    mkdir -p /root/.config/rclone/; \
    touch /root/.config/rclone/rclone.conf
