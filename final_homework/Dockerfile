FROM golang:1.18 AS builder

ARG APP_RELATIVE_PATH


COPY . /src
WORKDIR /src/app/${APP_RELATIVE_PATH}

RUN GOPROXY=https://goproxy.cn make build

FROM debian:stable-slim

ARG APP_RELATIVE_PATH

COPY --from=builder /src/app/${APP_RELATIVE_PATH}/bin /app

WORKDIR /app

VOLUME /data/conf

RUN if [ -f "./migrations" ] ; then ./migrations -conf /data/conf; fi
CMD ["./server", "-conf", "/data/conf"]
