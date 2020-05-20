FROM golang:1.14.3-alpine3.11 as builder

MAINTAINER smallnest <smallnest@gmail.com>

RUN echo "@community http://mirrors.ustc.edu.cn/alpine/edge/community" >> /etc/apk/repositories \
    && echo "@main http://mirrors.ustc.edu.cn/alpine/edge/main" >> /etc/apk/repositories \
    && sed -i 's/dl-cdn.alpinelinux.org/mirrors.ustc.edu.cn/g' /etc/apk/repositories \
    && apk add git \
    && apk update \
    bash git bash@main

ENV GOPROXY=https://goproxy.cn,direct
ENV GO111MODULE=on

RUN mkdir -p $GOPATH/src/github.com/smallnest \
    && cd $GOPATH/src/github.com/smallnest \
    && git clone --depth=1 https://github.com/smallnest/go-web-framework-benchmark.git \
    && cd $GOPATH/src/github.com/smallnest/go-web-framework-benchmark \
    && GO111MODULE=on go mod download \
    && go build -o  gowebbenchmark .

FROM alpine:3.11

MAINTAINER smallnest <smallnest@gmail.com>

RUN echo "@community http://mirrors.ustc.edu.cn/alpine/edge/community" >> /etc/apk/repositories \
    && echo "@main http://mirrors.ustc.edu.cn/alpine/edge/main" >> /etc/apk/repositories \
    && sed -i 's/dl-cdn.alpinelinux.org/mirrors.ustc.edu.cn/g' /etc/apk/repositories \
    && apk add \
    bash bash@main libressl3.0-libcrypto@main libressl3.0-libssl@main wrk@community gnuplot@community \
    ttf-dejavu ttf-droid ttf-freefont ttf-liberation ttf-ubuntu-font-family

VOLUME ["/data"]

COPY --from=builder /go/src/github.com/smallnest/go-web-framework-benchmark /go-web-framework-benchmark

WORKDIR /go-web-framework-benchmark

CMD ["/bin/sh","./docker-test.sh"]
