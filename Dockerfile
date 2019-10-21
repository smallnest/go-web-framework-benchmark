FROM golang:1.13.3-alpine3.10 as builder

MAINTAINER smallnest <smallnest@gmail.com>

RUN echo "@community http://dl-cdn.alpinelinux.org/alpine/edge/community" >> /etc/apk/repositories \
    && echo "@main http://dl-cdn.alpinelinux.org/alpine/edge/main" >> /etc/apk/repositories \
    && apk update \
    && apk add --no-cache --repository https://pkgs.alpinelinux.org/packages --allow-untrusted \
    bash git bash@main

ENV GOPROXY=https://goproxy.cn,direct
ENV GO111MODULE=on

RUN mkdir -p $GOPATH/src/github.com/smallnest \
    && cd $GOPATH/src/github.com/smallnest \
    && git clone https://github.com/smallnest/go-web-framework-benchmark.git \
    && cd $GOPATH/src/github.com/smallnest/go-web-framework-benchmark \
    && GO111MODULE=on go mod download \
    && go build -o  gowebbenchmark .

FROM alpine:3.10

MAINTAINER smallnest <smallnest@gmail.com>

RUN echo "@community http://dl-cdn.alpinelinux.org/alpine/edge/community" >> /etc/apk/repositories \
    && echo "@main http://dl-cdn.alpinelinux.org/alpine/edge/main" >> /etc/apk/repositories \
    && apk update \
    && apk add --no-cache --repository https://pkgs.alpinelinux.org/packages --allow-untrusted \
    bash bash@main libressl2.7-libcrypto@main libressl2.7-libssl@main wrk@community gnuplot@community \
    ttf-dejavu ttf-droid ttf-freefont ttf-liberation ttf-ubuntu-font-family

VOLUME ["/data"]

COPY --from=builder /go/src/github.com/smallnest/go-web-framework-benchmark /go-web-framework-benchmark

WORKDIR /go-web-framework-benchmark

CMD ["/bin/sh","./docker-test.sh"]
