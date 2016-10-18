FROM golang:1.7.1-alpine
MAINTAINER smallnest <smallnest@gmail.com>

RUN echo "@community http://dl-4.alpinelinux.org/alpine/edge/community" >> /etc/apk/repositories \
    && echo "@main http://dl-4.alpinelinux.org/alpine/edge/main" >> /etc/apk/repositories \
    && apk update \
    && apk add --no-cache --repository https://pkgs.alpinelinux.org/packages --allow-untrusted \
    bash git libressl2.4-libcrypto@main libressl2.4-libssl@main wrk@community gnuplot@community

# RUN go get github.com/smallnest/go-web-framework-benchmark \
#     && cd $GOPATH/src/github.com/smallnest/go-web-framework-benchmark \
#     && go build -o  gowebbenchmark server.go

RUN go get github.com/smallnest/go-web-framework-benchmark && mkdir /data

VOLUME ["/data"]

WORKDIR $GOPATH/src/github.com/smallnest/go-web-framework-benchmark

CMD ["/bin/sh","./docker-test.sh"]
