#!/bin/bash

echo "Running test.sh" > test.log
./test.sh
./testresults/plot.sh
echo "Finished test.sh" >> test.log

echo "Running test-latency.sh" >> test.log
./test-latency.sh
./testresults/plot.sh
echo "Finished test-latency.sh" >> test.log

echo "Running test-pipelining.sh" >> test.log
./test-pipelining.sh
./testresults/plot.sh
echo "Finished test-pipelining.sh" >> test.log
