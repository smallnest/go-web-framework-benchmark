# go-web-framework-benchmark
This benchmark suite aims to compare the performance of Go web frameworks. It is inspired by [Go HTTP Router Benchmark](https://github.com/julienschmidt/go-http-routing-benchmark) but this benchmark suite is different with that. Go HTTP Router Benchmark suit aims to compare the performance of **routers** but this Benchmark suit aims to compare whole HTTP request processing.

**Last Test Updated:** 2020-05

*test environment*

* CPU:      QEMU Virtual CPU version(1.8GHZ, 8 cores)
* Memory:   16G
* Go:       go1.16.3 linux/amd64
* OS:       CentOS Linux release 7.5.1804 (Core)


## Tested web frameworks (in alphabetical order)

**Only test those webframeworks which are stable**

* [ace](https://github.com/plimble/ace)
* [atreugo](https://github.com/savsgio/atreugo)
* [aurora](https://github.com/aurora-go/aurora)
* [baa](https://github.com/go-baa/baa)
* [beego](https://github.com/astaxie/beego)
* [bone](https://github.com/go-zoo/bone)
* [chi](https://github.com/go-chi/chi)
* [clevergo](https://github.com/clevergo/clevergo)
* [default http](https://golang.org/pkg/net/http/)
* [denco](https://github.com/naoina/denco)
* [echo](https://github.com/labstack/echo)
* [fasthttp-routing](https://github.com/qiangxue/fasthttp-routing)
* [fasthttp/router](https://github.com/fasthttp/router)
* [fasthttp](https://github.com/valyala/fasthttp)
* [fasthttprouter](https://github.com/buaazp/fasthttprouter)
* [fastRouter](https://github.com/razonyang/fastrouter)
* [fiber](https://gofiber.io/)
* [flygo](https://github.com/billcoding/flygo)
* [gear](http://github.com/teambition/gear)
* [gearbox](https://github.com/gogearbox/gearbox)
* [gem](https://github.com/go-gem/gem)
* [gin](https://github.com/gin-gonic/gin)
* [goframe](https://github.com/gogf/gf)
* [go-json-rest](https://github.com/ant0ine/go-json-rest)
* [go-ozzo](https://github.com/go-ozzo/ozzo-routing)
* [go-restful](https://github.com/emicklei/go-restful)
* [go-tigertonic](https://github.com/rcrowley/go-tigertonic)
* [goji](https://github.com/zenazn/goji/web)
* [goji](http://goji.io)
* [golf](https://github.com/dinever/golf)
* [gongular](https://github.com/mustafaakin/gongular)
* [gorilla](https://github.com/gorilla/mux)
* [gorouter](https://github.com/vardius/gorouter)
* [goyave](https://github.com/System-Glitch/goyave)
* [~~gramework~~](https://github.com/gramework/gramework):its init function pollutes the test so it has been removed from test.
* [httprouter](https://github.com/julienschmidt/httprouter)
* [httptreemux](https://github.com/dimfeld/httptreemux)
* [lars](https://github.com/go-playground/lars)
* [lion](https://github.com/celrenheit/lion)
* [macaron](https://github.com/Unknwon/macaron)
* [martini](https://github.com/go-martini/martini)
* [muxie](https://github.com/kataras/muxie)
* [negroni](https://github.com/urfave/negroni)
* [neo](https://github.com/ivpusic/neo)
* [pat](https://github.com/bmizerany/pat)
* [pure](https://github.com/go-playground/pure)
* [r2router](https://github.com/vanng822/r2router)
* [siris](https://github.com/go-siris/siris)
* [tango](https://github.com/lunny/tango)
* [tinyrouter](https://github.com/go101/tinyrouter)
* [traffic](https://github.com/pilu/traffic)
* [treemux](https://github.com/vmihailenco/treemux)
* [violetear](https://github.com/nbari/violetear)
* [vulcan](https://github.com/mailgun/route)
* [webgo](https://github.com/bnkamalesh/webgo)

**some libs have not been maintained and the test code has removed them**

## Motivation
When I investigated performance of Go web frameworks, I found [Go HTTP Router Benchmark](https://github.com/julienschmidt/go-http-routing-benchmark), created by Julien Schmidt. He also developed a high performance http router: [httprouter](https://github.com/julienschmidt/httprouter). I had thought I got the performance result until I created a piece of codes to mock the real business logics:

```go
api.Get("/rest/hello", func(c *XXXXX.Context) {
	sleepTime := strconv.Atoi(os.Args[1]) //10ms
	if sleepTime > 0 {
		time.Sleep(time.Duration(sleepTime) * time.Millisecond)
	}

	c.Text("Hello world")
})
```

When I use the above codes to test those web frameworks, the token time of route selection is not so important in the whole http request processing, although performance of route selection of web frameworks are very different.

So I create this project to compare performance of web frameworks including connection, route selection, handler processing. It mocks business logics and can set a special processing time.

The you can get some interesting results if you use it to test.

## Implementation
When you test a web framework, this test suit will starts a simple http server implemented by this web framework. It is a real http server and only contains GET url: "/hello".

When this server processes this url, it will sleep n milliseconds in this handler. It mocks the business logics such as:
* read data from sockets
* write data to disk
* access databases
* access cache servers
* invoke other microservices
* ……

It contains a test.sh that can do those tests automatically.

It uses [wrk](https://github.com/wg/wrk/) to test.

## Basic Test
The first test case is to mock 0 ms, 10 ms, 100 ms, 500 ms processing time in handlers.

![Benchmark (Round 3)](benchmark.png)
the concurrency clients are 5000.

![Latency (Round 3)](benchmark_latency.png)
Latency is the time of real processing time by web servers. The smaller is the better.

![Allocs (Round 3)](benchmark_alloc.png)
Allocs is the heap allocations by web servers when test is running. The unit is MB. The smaller is the better.


If we enable http pipelining, test result as below:

![benchmark pipelining (Round 2)](benchmark-pipeline.png)

## Concurrency Test
In 30 ms processing time, the test result for 100, 1000, 5000 clients is:

![concurrency (Round 3)](concurrency.png)

![Latency (Round 3)](concurrency_latency.png)

![Latency (Round 3)](concurrency_alloc.png)


If we enable http pipelining, test result as below:

![concurrency pipelining(Round 2)](concurrency-pipeline.png)


## cpu-bound case Test

![cpu-bound (5000 concurrency)](cpubound_benchmark.png)

## Usage
You should install this package first if you want to run this test.

```
go get github.com/smallnest/go-web-framework-benchmark
```

It takes a while to install a large number of dependencies that need to be downloaded. Once that command completes, you can run:

```
cd $GOPATH/src/github.com/smallnest/go-web-framework-benchmark
go build -o  gowebbenchmark *.go
./test.sh
```

It will  generate test results in processtime.csv and concurrency.csv. You can modify test.sh to execute your customized test cases.


* If you also want to generate latency data and allocation data, you can run the script:
```
./test-latency.sh
```

* If you don't want use keepalive, you can run:
```
./test-latency-nonkeepalive.sh
```

* If you want to test http pipelining, you can run:
```
./test-pipelining.sh
```

* If you want to test some of web frameworks, you can modify the test script and only keep your selected web frameworks:
```
……
web_frameworks=( "default" "ace" "beego" "bone" "denco" "echov1" "echov2standard" "echov2fasthttp" "fasthttp-raw" "fasthttprouter" "fasthttp-routing" "gin" "gocraftWeb" "goji" "gojiv2" "gojsonrest" "gorestful" "gorilla" "httprouter" "httptreemux" "lars" "lion" "macaron" "martini" "pat" "r2router" "tango" "tiger" "traffic" "violetear" "vulcan")
……
```
* If you want to test all cases, you can run:

```
./test-all.sh
```

## Plot
you can run the shell script `plot.sh` in testresults directory and it can generate all images in its parent directory.


### Add new web framework
Welcome to add new Go web frameworks. You can follow the below steps and send me a pull request.

1. add your web framework link in README
2. add a hello implementation in server.go
3. add your webframework in libs.sh

Please add your web framework alphabetically.
