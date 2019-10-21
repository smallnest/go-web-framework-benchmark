#!/bin/bash

chmod +x *.sh

./test-latency.sh
./test-pipelining.sh
cd testresults
./plot.sh
cd ..
cp -R testresults /data/testresults
cp *.png /data/testresults