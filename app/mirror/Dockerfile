FROM golang:1.22-bookworm AS builder
WORKDIR /src
COPY . .
RUN set -e; \
    go build -o ./bin/mirror -mod vendor -ldflags "-w -s" -trimpath ./app/mirror

FROM debian:12
ENV MIRROR_SERVER_BIND=0.0.0.0:80 \
    MIRROR_SYSTEM_DATA=/data \
    DEBIAN_FRONTEND=noninteractive
RUN set -e; \
    mkdir -p /etc/apt/sources.list.d/; \
    rm -rf /etc/apt/sources.list.d/debian.sources; \
    echo "\
deb http://mirrors.ustc.edu.cn/debian/ bookworm main contrib non-free non-free-firmware\n\
deb http://mirrors.ustc.edu.cn/debian/ bookworm-updates main contrib non-free non-free-firmware\n\
deb http://mirrors.ustc.edu.cn/debian/ bookworm-backports main contrib non-free non-free-firmware\n\
deb http://mirrors.ustc.edu.cn/debian-security bookworm-security main contrib non-free non-free-firmware\n" > /etc/apt/sources.list.d/ustc.list; \
    apt-get update; \
    apt-get install -y apt-transport-https ca-certificates; \
    echo "\
deb https://mirrors.ustc.edu.cn/debian/ bookworm main contrib non-free non-free-firmware\n\
deb https://mirrors.ustc.edu.cn/debian/ bookworm-updates main contrib non-free non-free-firmware\n\
deb https://mirrors.ustc.edu.cn/debian/ bookworm-backports main contrib non-free non-free-firmware\n\
deb https://mirrors.ustc.edu.cn/debian-security bookworm-security main contrib non-free non-free-firmware\n" > /etc/apt/sources.list.d/ustc.list; \
    apt-get update; \
    apt-get upgrade -y; \
    apt-get install -y apt-transport-https ca-certificates
COPY --from=builder /src/bin/mirror /usr/bin/mirror-server
COPY --from=builder /src/data/ /data/
EXPOSE 80
ENTRYPOINT ["mirror-server"]