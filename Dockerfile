FROM golang:1.9.0-alpine
MAINTAINER smallnest <smallnest@gmail.com>

RUN echo "@community http://dl-4.alpinelinux.org/alpine/edge/community" >> /etc/apk/repositories \
    && echo "@main http://dl-4.alpinelinux.org/alpine/edge/main" >> /etc/apk/repositories \
    && apk update \
    && apk add --no-cache --repository https://pkgs.alpinelinux.org/packages --allow-untrusted \
    bash git bash@main libressl2.5-libcrypto@main libressl2.5-libssl@main wrk@community gnuplot@community \
    ttf-dejavu ttf-droid ttf-freefont ttf-liberation ttf-ubuntu-font-family

# RUN go get github.com/smallnest/go-web-framework-benchmark \
#     && cd $GOPATH/src/github.com/smallnest/go-web-framework-benchmark \
#     && go build -o  gowebbenchmark server.go

# RUN go get github.com/smallnest/go-web-framework-benchmark && mkdir /data


VOLUME ["/data"]

# add current version of gowebbenchmark
RUN mkdir -p $GOPATH/src/github.com/smallnest/go-web-framework-benchmark
ADD . $GOPATH/src/github.com/smallnest/go-web-framework-benchmark/

WORKDIR $GOPATH/src/github.com/smallnest/go-web-framework-benchmark

CMD ["/bin/sh","./docker-test.sh"]
