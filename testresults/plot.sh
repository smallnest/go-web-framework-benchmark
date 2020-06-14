#!/bin/bash
m_path=$(dirname $0)
m_path=${m_path/\./$(pwd)}
cd $m_path

./transpose.sh

gnuplot  benchmark.gnu
gnuplot  benchmark_latency.gnu
gnuplot  benchmark_alloc.gnu
gnuplot  benchmark_pipeline.gnu

gnuplot  concurrency.gnu
gnuplot  concurrency_latency.gnu
gnuplot  concurrency_alloc.gnu
gnuplot  concurrency_pipeline.gnu

gnuplot  cpubound_benchmark.gnu
gnuplot  cpubound_concurrency.gnu

rm -fr t_*.csv
