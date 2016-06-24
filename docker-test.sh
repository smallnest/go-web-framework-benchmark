#!/bin/bash

go get -u github.com/smallnest/go-web-framework-benchmark
go build -o gowebbenchmark server.go

chmod +x *.sh

./test-latency.sh
./test-pipelining.sh
cp -R testresults /data/testresults