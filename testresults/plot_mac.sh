#!/bin/bash
m_path=$(dirname $0)
m_path=${m_path/\./$(pwd)}
cd $m_path

./transpose.sh

gnuplot -c benchmark.gnu
gnuplot -c benchmark_latency.gnu
gnuplot -c benchmark_alloc.gnu
gnuplot -c benchmark_pipeline.gnu

gnuplot -c concurrency.gnu
gnuplot -c concurrency_latency.gnu
gnuplot -c concurrency_alloc.gnu
gnuplot -c concurrency_pipeline.gnu

gnuplot -c cpubound_benchmark.gnu
gnuplot -c cpubound_concurrency.gnu

rm -fr t_*.csv
