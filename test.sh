#!/bin/bash

server_bin_name="gowebbenchmark"

web_frameworks=( "default" "ace" "beego" "bone" "denco" "echov1" "echov2standard" "echov2fasthttp" "fasthttp-raw" "fasthttprouter" "fasthttp-routing" "gin" "gocraftWeb" "goji" "gojiv2" "gojsonrest" "gorestful" "gorilla" "httprouter" "httptreemux" "iris" "lars" "lion" "macaron" "martini" "pat" "possum" "r2router" "tango" "tiger" "traffic" "vulcan" )
length=${#web_frameworks[@]}

test_result=()

test_web_framework()
{
  echo "testing web framework: $2"
  ./$server_bin_name $2 $3 &
  sleep 5

  throughput=`wrk -t16 -c$4 -d30s http://127.0.0.1:8080/hello | grep Requests/sec | awk '{print $2}'`
  echo "throughput: $throughput requests/second"
  test_result[$1]=$throughput

  pkill -9 $server_bin_name
  sleep 5
  echo "finsihed testing $2"
  echo
}

test_all()
{
  echo "###################################"
  echo "                                   "
  echo "      ProcessingTime  $1ms         "
  echo "      Concurrency     $2           "
  echo "                                   "
  echo "###################################"
  for ((i=0; i<$length; i++))
  do
  	test_web_framework $i ${web_frameworks[$i]} $1 $2
  done
}


pkill -9 $server_bin_name

echo ","$(IFS=$','; echo "${web_frameworks[*]}" ) > processtime.csv
test_all 0 5000
echo "0 ms,"$(IFS=$','; echo "${test_result[*]}" ) >> processtime.csv
test_all 10 5000
echo "10 ms,"$(IFS=$','; echo "${test_result[*]}" ) >> processtime.csv
test_all 100 5000
echo "100 ms,"$(IFS=$','; echo "${test_result[*]}" ) >> processtime.csv
test_all 500 5000
echo "500 ms,"$(IFS=$','; echo "${test_result[*]}" ) >> processtime.csv



echo ","$(IFS=$','; echo "${web_frameworks[*]}" ) > concurrency.csv
test_all 30 100
echo "100,"$(IFS=$','; echo "${test_result[*]}" ) >> concurrency.csv
test_all 30 1000
echo "1000,"$(IFS=$','; echo "${test_result[*]}" ) >> concurrency.csv
test_all 30 5000
echo "5000,"$(IFS=$','; echo "${test_result[*]}" ) >> concurrency.csv
