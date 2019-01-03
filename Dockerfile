FROM golang:1.11.4-alpine3.7
MAINTAINER smallnest <smallnest@gmail.com>

RUN echo "@community http://dl-cdn.alpinelinux.org/alpine/edge/community" >> /etc/apk/repositories \
    && echo "@main http://dl-cdn.alpinelinux.org/alpine/edge/main" >> /etc/apk/repositories \
    && apk update \
    && apk add --no-cache --repository https://pkgs.alpinelinux.org/packages --allow-untrusted \
    bash git bash@main libressl2.7-libcrypto@main libressl2.7-libssl@main wrk@community gnuplot@community \
    ttf-dejavu ttf-droid ttf-freefont ttf-liberation ttf-ubuntu-font-family

RUN go get github.com/smallnest/go-web-framework-benchmark \
    && cd $GOPATH/src/github.com/smallnest/go-web-framework-benchmark \
    && go build -o  gowebbenchmark server.go

VOLUME ["/data"]


WORKDIR $GOPATH/src/github.com/smallnest/go-web-framework-benchmark

CMD ["/bin/sh","./docker-test.sh"]
