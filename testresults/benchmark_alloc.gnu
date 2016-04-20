reset
set title "Benchmark of different processing time (Allocations)"
set boxwidth 0.9
set datafile separator ","
set style data histogram
set style histogram clustered gap 2
set style fill solid 0.7 border
set border lw 0.8

set ylabel "Allocations (MB)"
set xtics nomirror rotate by -45
set ytics nomirror

set border 1+2 back 
set boxwidth -2

set grid

set term pngcairo font "Times Roman,14"  enhanced size 1024,600 background rgb "gray80"
set output "../benchmark_alloc.png"

plot 't_processtime_alloc.csv' using 2:xticlabels(1) title columnheader(2), '' using 3:xticlabels(1) title columnheader(3), '' using 4:xticlabels(1) title columnheader(4), '' using 5:xticlabels(1) title columnheader(5)