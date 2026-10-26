FROM ghcr.io/kopeisec/debian:13 AS builder
WORKDIR /src
COPY ./build/ ./
RUN \
    case $(uname -m) in \
    x86_64)  cp mirror_linux_amd64 mirror ;; \
    aarch64) cp mirror_linux_arm64 mirror ;; \
    *) echo "[ ERROR ] unsupported platform: $(uname -m)" && false ;; \
    esac

FROM ghcr.io/kopeisec/debian:13
ENV MIRROR_SERVER_BIND=0.0.0.0:80 \
    MIRROR_SYSTEM_DATA=/data \
    DEBIAN_FRONTEND=noninteractive
COPY --from=builder /src/mirror /usr/bin/mirror-server
COPY ./data/ /data/
EXPOSE 80
ENTRYPOINT ["mirror-server"]