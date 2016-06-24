FROM golang:1.6.2-alpine
MAINTAINER smallnest <smallnest@gmail.com>

RUN echo "@testing http://dl-4.alpinelinux.org/alpine/edge/testing" >> /etc/apk/repositories \
    && apk update \
    && apk add --no-cache --repository https://pkgs.alpinelinux.org/packages --allow-untrusted \
    bash git wrk@testing gnuplot@testing

# RUN go get github.com/smallnest/go-web-framework-benchmark \
#     && cd $GOPATH/src/github.com/smallnest/go-web-framework-benchmark \
#     && go build -o  gowebbenchmark server.go

RUN go get github.com/smallnest/go-web-framework-benchmark && mkdir /data

VOLUME ["/data"]

WORKDIR $GOPATH/src/github.com/smallnest/go-web-framework-benchmark

CMD ["/bin/sh","$GOPATH/src/github.com/smallnest/go-web-framework-benchmark/docker-test.sh"]
