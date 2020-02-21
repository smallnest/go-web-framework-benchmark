package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"runtime"
	"strconv"
	"time"

	"github.com/ant0ine/go-json-rest/rest"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"github.com/bmizerany/pat"
	"github.com/bnkamalesh/webgo"
	"github.com/buaazp/fasthttprouter"
	"github.com/dimfeld/httptreemux"
	"github.com/dinever/golf"
	restful "github.com/emicklei/go-restful"
	fasthttpSlashRouter "github.com/fasthttp/router"
	"github.com/gofiber/fiber"
	"github.com/gin-gonic/gin"
	"github.com/gramework/gramework"
	"github.com/kataras/muxie"
	"github.com/savsgio/atreugo/v10"

	// "github.com/go-siris/siris"
	// siriscontext "github.com/go-siris/siris/context"
	"github.com/nbari/violetear"
	"github.com/urfave/negroni"
	macaron "gopkg.in/macaron.v1"

	// "github.com/go-gas/gas" // NOTE(@kirilldanshin): gas is 404 now, comment out
	bxog "github.com/claygod/Bxog"
	"github.com/go-martini/martini"
	ozzo "github.com/go-ozzo/ozzo-routing"
	"github.com/go-playground/lars"
	"github.com/go-playground/pure"
	"github.com/go-zoo/bone"
	"github.com/gocraft/web"
	"github.com/gorilla/mux"
	gowwwrouter "github.com/gowww/router"
	tiny "go101.org/tinyrouter"

	"github.com/ivpusic/neo"
	"github.com/julienschmidt/httprouter"
	echo "github.com/labstack/echo/v4"
	llog "github.com/lunny/log"
	"github.com/lunny/tango"
	vulcan "github.com/mailgun/route"
	"github.com/mustafaakin/gongular"
	"github.com/naoina/denco"
	"github.com/pilu/traffic"

	// "github.com/plimble/ace"
	"github.com/pressly/chi"
	routing "github.com/qiangxue/fasthttp-routing"
	"github.com/razonyang/fastrouter"
	tigertonic "github.com/rcrowley/go-tigertonic"
	"github.com/teambition/gear"
	"github.com/tockins/fresh"
	"github.com/valyala/fasthttp"
	"github.com/vanng822/r2router"
	"github.com/vardius/gorouter/v4"
	goji "goji.io"
	gojipat "goji.io/pat"
	gcontext "golang.org/x/net/context"
	baa "gopkg.in/baa.v1"
	lion "gopkg.in/celrenheit/lion.v1"
)

var port = 8080
var sleepTime = 0
var cpuBound bool
var target = 15
var sleepTimeDuration time.Duration
var message = []byte("hello world")
var messageStr = "hello world"
var samplingPoint = 20 //seconds

// server [default] [10] [8080]
func main() {
	args := os.Args
	argsLen := len(args)
	webFramework := "default"
	if argsLen > 1 {
		webFramework = args[1]
	}
	if argsLen > 2 {
		sleepTime, _ = strconv.Atoi(args[2])
		if sleepTime == -1 {
			cpuBound = true
			sleepTime = 0
		}
	}
	if argsLen > 3 {
		port, _ = strconv.Atoi(args[3])
	}
	if argsLen > 4 {
		samplingPoint, _ = strconv.Atoi(args[4])
	}
	sleepTimeDuration = time.Duration(sleepTime) * time.Millisecond
	samplingPointDuration := time.Duration(samplingPoint) * time.Second

	go func() {
		time.Sleep(samplingPointDuration)
		var mem runtime.MemStats
		runtime.ReadMemStats(&mem)
		var u uint64 = 1024 * 1024
		fmt.Printf("TotalAlloc: %d\n", mem.TotalAlloc/u)
		fmt.Printf("Alloc: %d\n", mem.Alloc/u)
		fmt.Printf("HeapAlloc: %d\n", mem.HeapAlloc/u)
		fmt.Printf("HeapSys: %d\n", mem.HeapSys/u)
	}()

	switch webFramework {
	case "default":
		startDefaultMux()
	// case "ace":
	// 	startAce()
	case "atreugo":
		startAtreugo()
	case "baa":
		startBaa()
	case "beego":
		startBeego()
	case "bone":
		startBone()
	case "bxog":
		startBxog()
	case "chi":
		startChi()
	case "denco":
		startDenco()
	case "echo":
		startEcho()
	case "fasthttp-raw":
		startFasthttp()
	case "fasthttprouter":
		startFastHTTPRouter()
	case "fasthttp/router":
		startFastHTTPSlashRouter()
	case "fasthttp-routing":
		startFastHTTPRouting()
	case "fastrouter":
		startFastRouter()
	case "fiber":
		startFiber()
	case "fresh":
		startFresh()
	case "gear":
		startGear()
	case "gin":
		startGin()
	case "gocraftWeb":
		startGocraftWeb()
	case "goji":
		startGoji()
	case "gojsonrest":
		startGoJSONRest()
	case "golf":
		startGolf()
	case "gongular":
		startGongular()
	case "gorestful":
		startGoRestful()
	case "gorilla":
		startGorilla()
	case "gorouter":
		startGorouter()
	case "gorouterfasthttp":
		startGorouterFastHTTP()
	case "go-ozzo":
		startGoozzo()
	case "gowww":
		startGowww()
	case "gramework":
		startGramework()
	case "httprouter":
		startHTTPRouter()
	case "httptreemux":
		starthttpTreeMux()
	case "lars":
		startLars()
	case "lion":
		startLion()
	case "macaron":
		startMacaron()
	case "martini":
		startMartini()
	case "muxie":
		startMuxie()
	case "negroni":
		startNegroni()
	case "neo":
		startNeo()
	case "pat":
		startPat()
	case "pure":
		startPure()
	case "r2router":
		startR2router()
	// case "siris":
	// 	startSirisrouter()
	case "tango":
		startTango()
	case "tiger":
		startTigerTonic()
	case "tinyrouter":
		startTinyRouter()
	case "traffic":
		startTraffic()
	case "violetear":
		startVioletear()
	case "vulcan":
		startVulcan()
	case "webgo":
		startWebgo()
	}
}

// default mux
func helloHandler(w http.ResponseWriter, r *http.Request) {
	if cpuBound {
		pow(target)
	} else {

		if sleepTime > 0 {
			time.Sleep(sleepTimeDuration)
		} else {
			runtime.Gosched()
		}
	}
	w.Write(message)
}
func startDefaultMux() {
	http.HandleFunc("/hello", helloHandler)
	http.ListenAndServe(":"+strconv.Itoa(port), nil)
}

// ace
// func aceHandler(c *ace.C) {
// 	if sleepTime > 0 {
// 		time.Sleep(sleepTimeDuration)
// 	} else {
// 		runtime.Gosched()
// 	}
// 	c.Writer.Write(message)
// }
// func startAce() {
// 	mux := ace.New()
// 	mux.GET("/hello", aceHandler)
// 	mux.Run(":" + strconv.Itoa(port))
// }

// atreugo
func atreugoHandler(ctx *atreugo.RequestCtx) error {
	if cpuBound {
		pow(target)
	} else {

		if sleepTime > 0 {
			time.Sleep(sleepTimeDuration)
		} else {
			runtime.Gosched()
		}
	}
	return ctx.TextResponse(messageStr)
}

func startAtreugo() {
	mux := atreugo.New(&atreugo.Config{Addr: "127.0.0.1:" + strconv.Itoa(port)})
	mux.GET("/hello", atreugoHandler)
	mux.ListenAndServe()
}

// baa
func baaHandler(ctx *baa.Context) {
	if cpuBound {
		pow(target)
	} else {

		if sleepTime > 0 {
			time.Sleep(sleepTimeDuration)
		} else {
			runtime.Gosched()
		}
	}
	ctx.Text(200, message)
}
func startBaa() {
	mux := baa.New()
	mux.Get("/hello", baaHandler)
	mux.Run(":" + strconv.Itoa(port))
}

//beego
func beegoHandler(ctx *context.Context) {
	if cpuBound {
		pow(target)
	} else {

		if sleepTime > 0 {
			time.Sleep(sleepTimeDuration)
		} else {
			runtime.Gosched()
		}
	}
	ctx.WriteString(messageStr)
}
func startBeego() {
	beego.BConfig.RunMode = beego.PROD
	beego.BeeLogger.Close()
	mux := beego.NewControllerRegister()
	mux.Get("/hello", beegoHandler)
	http.ListenAndServe(":"+strconv.Itoa(port), mux)
}

// bone
func startBone() {
	mux := bone.New()
	mux.HandleFunc("/hello", helloHandler)
	http.ListenAndServe(":"+strconv.Itoa(port), mux)
}

// bxog
func bxogHandler(w http.ResponseWriter, req *http.Request, r *bxog.Router) {
	if cpuBound {
		pow(target)
	} else {

		if sleepTime > 0 {
			time.Sleep(sleepTimeDuration)
		} else {
			runtime.Gosched()
		}
	}
	w.Write(message)
}
func startBxog() {
	mux := bxog.New()
	mux.Add("/hello", bxogHandler)
	mux.Start(":" + strconv.Itoa(port))
}

//chi
func startChi() {
	// Create a router instance.
	r := chi.NewRouter()

	// Register route handler.
	r.Get("/hello", helloHandler)

	// Start Chi.
	http.ListenAndServe(":"+strconv.Itoa(port), r)
}

// denco
func dencoHandler(w http.ResponseWriter, r *http.Request, params denco.Params) {
	if cpuBound {
		pow(target)
	} else {

		if sleepTime > 0 {
			time.Sleep(sleepTimeDuration)
		} else {
			runtime.Gosched()
		}
	}
	w.Write(message)
}
func startDenco() {
	mux := denco.NewMux()
	handler, _ := mux.Build([]denco.Handler{mux.GET("/hello", denco.HandlerFunc(dencoHandler))})
	http.ListenAndServe(":"+strconv.Itoa(port), handler)
}

// echo
func echoHandler(c echo.Context) error {
	if cpuBound {
		pow(target)
	} else {

		if sleepTime > 0 {
			time.Sleep(sleepTimeDuration)
		} else {
			runtime.Gosched()
		}
	}
	c.Response().Write(message)
	return nil
}
func startEcho() {
	e := echo.New()
	e.GET("/hello", echoHandler)

	e.Start(":" + strconv.Itoa(port))
}

//fasthttp
func fastHTTPRawHandler(ctx *fasthttp.RequestCtx) {
	if string(ctx.Method()) == "GET" {
		switch string(ctx.Path()) {
		case "/hello":
			if cpuBound {
				pow(target)
			} else {

				if sleepTime > 0 {
					time.Sleep(sleepTimeDuration)
				} else {
					runtime.Gosched()
				}
			}
			ctx.Write(message)
		default:
			ctx.Error("Unsupported path", fasthttp.StatusNotFound)
		}
		return
	}
	ctx.Error("Unsupported method", fasthttp.StatusMethodNotAllowed)
}
func startFasthttp() {
	fasthttp.ListenAndServe(":"+strconv.Itoa(port), fastHTTPRawHandler)
}

//fasthttprouter
func fastHTTPHandler(ctx *fasthttp.RequestCtx) {
	if cpuBound {
		pow(target)
	} else {

		if sleepTime > 0 {
			time.Sleep(sleepTimeDuration)
		} else {
			runtime.Gosched()
		}
	}
	ctx.Write(message)
}
func startFastHTTPRouter() {
	mux := fasthttprouter.New()
	mux.GET("/hello", fastHTTPHandler)
	fasthttp.ListenAndServe(":"+strconv.Itoa(port), mux.Handler)
}

//fasthttp Router
func startFastHTTPSlashRouter() {
	mux := fasthttpSlashRouter.New()
	mux.GET("/hello", fastHTTPHandler)
	fasthttp.ListenAndServe(":"+strconv.Itoa(port), mux.Handler)
}

//fasthttprouting
func fastHTTPRoutingHandler(c *routing.Context) error {
	if cpuBound {
		pow(target)
	} else {

		if sleepTime > 0 {
			time.Sleep(sleepTimeDuration)
		} else {
			runtime.Gosched()
		}
	}
	c.Write(message)
	return nil
}
func startFastHTTPRouting() {
	mux := routing.New()
	mux.Get("/hello", fastHTTPRoutingHandler)
	fasthttp.ListenAndServe(":"+strconv.Itoa(port), mux.HandleRequest)
}

//fastrouter
func fastRouterHandler(w http.ResponseWriter, r *http.Request) {
	if cpuBound {
		pow(target)
	} else {

		if sleepTime > 0 {
			time.Sleep(sleepTimeDuration)
		} else {
			runtime.Gosched()
		}
	}
	w.Write(message)
}
func startFastRouter() {
	mux := fastrouter.New()
	mux.Get("/hello", fastRouterHandler)
	mux.Prepare()
	http.ListenAndServe(":"+strconv.Itoa(port), mux)
}

//fresh
func freshHandler(c fresh.Context) error {
	if cpuBound {
		pow(target)
	} else {

		if sleepTime > 0 {
			time.Sleep(sleepTimeDuration)
		} else {
			runtime.Gosched()
		}
	}
	c.Response().Text(http.StatusOK, messageStr)
	return nil
}

func startFresh() {
	f := fresh.New()
	f.Config().Port = port
	f.GET("/hello", freshHandler)
	f.Start()
}

// fiber
func fiberHandler(ctx *fiber.Ctx) {
	if cpuBound {
		pow(target)
	} else {

		if sleepTime > 0 {
			time.Sleep(sleepTimeDuration)
		} else {
			runtime.Gosched()
		}
	}
	ctx.SendString(messageStr)
}

func startFiber() {
	app := fiber.New()
	app.Get("/hello", fiberHandler)
	app.Listen(port)
}

//gear
func startGear() {
	app := gear.New()
	router := gear.NewRouter()

	router.Get("/hello", func(c *gear.Context) error {
		if cpuBound {
			pow(target)
		} else {

			if sleepTime > 0 {
				time.Sleep(sleepTimeDuration)
			} else {
				runtime.Gosched()
			}
		}
		return c.HTML(200, messageStr)
	})
	app.UseHandler(router)
	app.Listen(":" + strconv.Itoa(port))
}

// gin
func ginHandler(c *gin.Context) {
	if cpuBound {
		pow(target)
	} else {

		if sleepTime > 0 {
			time.Sleep(sleepTimeDuration)
		} else {
			runtime.Gosched()
		}
	}
	c.Writer.Write(message)
}
func startGin() {
	gin.SetMode(gin.ReleaseMode)
	mux := gin.New()
	mux.GET("/hello", ginHandler)
	mux.Run(":" + strconv.Itoa(port))
}

// gocraftWeb
type gocraftWebContext struct{}

func gocraftWebHandler(w web.ResponseWriter, r *web.Request) {
	if cpuBound {
		pow(target)
	} else {

		if sleepTime > 0 {
			time.Sleep(sleepTimeDuration)
		} else {
			runtime.Gosched()
		}
	}
	w.Write(message)
}
func startGocraftWeb() {
	mux := web.New(gocraftWebContext{})
	mux.Get("/hello", gocraftWebHandler)
	http.ListenAndServe(":"+strconv.Itoa(port), mux)
}

// goji
func gojiHandler(w http.ResponseWriter, r *http.Request) {
	if cpuBound {
		pow(target)
	} else {

		if sleepTime > 0 {
			time.Sleep(sleepTimeDuration)
		} else {
			runtime.Gosched()
		}
	}
	w.Write(message)
}
func startGoji() {
	mux := goji.NewMux()
	mux.HandleFunc(gojipat.Get("/hello"), gojiHandler)
	http.ListenAndServe(":"+strconv.Itoa(port), mux)
}

// goJsonRest
func goJSONRestHandler(w rest.ResponseWriter, req *rest.Request) {
	if cpuBound {
		pow(target)
	} else {

		if sleepTime > 0 {
			time.Sleep(sleepTimeDuration)
		} else {
			runtime.Gosched()
		}
	}
	iow := w.(io.Writer)
	iow.Write(message)
}
func startGoJSONRest() {
	api := rest.NewApi()
	router, _ := rest.MakeRouter(
		&rest.Route{HttpMethod: "GET", PathExp: "/hello", Func: goJSONRestHandler},
	)
	api.SetApp(router)
	http.ListenAndServe(":"+strconv.Itoa(port), api.MakeHandler())
}

func golfHandler(ctx *golf.Context) {
	if cpuBound {
		pow(target)
	} else {

		if sleepTime > 0 {
			time.Sleep(sleepTimeDuration)
		} else {
			runtime.Gosched()
		}
	}
	ctx.Send(messageStr)
}

func startGolf() {
	app := golf.New()
	app.Get("/hello", golfHandler)
	app.Run(":" + strconv.Itoa(port))
}

type HelloMessage struct{}

func (w *HelloMessage) Handle(c *gongular.Context) error {
	if cpuBound {
		pow(target)
	} else {

		if sleepTime > 0 {
			time.Sleep(sleepTimeDuration)
		} else {
			runtime.Gosched()
		}
	}
	c.SetBody(messageStr)

	return nil
}

func startGongular() {
	g := gongular.NewEngine()
	g.GetRouter().GET("/hello", &HelloMessage{})
	g.ListenAndServe(":" + strconv.Itoa(port))
}

// goRestful
func goRestfulHandler(r *restful.Request, w *restful.Response) {
	if cpuBound {
		pow(target)
	} else {

		if sleepTime > 0 {
			time.Sleep(sleepTimeDuration)
		} else {
			runtime.Gosched()
		}
	}
	w.Write(message)
}
func startGoRestful() {
	wsContainer := restful.NewContainer()
	ws := new(restful.WebService)
	ws.Route(ws.GET("/hello").To(goRestfulHandler))
	wsContainer.Add(ws)
	http.ListenAndServe(":"+strconv.Itoa(port), wsContainer)
}

// gorilla
func startGorilla() {
	mux := mux.NewRouter()
	mux.HandleFunc("/hello", helloHandler).Methods("GET")
	http.ListenAndServe(":"+strconv.Itoa(port), mux)
}

// gorouter
func startGorouter() {
	router := gorouter.New()
	router.GET("/hello", http.HandlerFunc(helloHandler))
	http.ListenAndServe(":"+strconv.Itoa(port), router)
}

func startGorouterFastHTTP() {
	router := gorouter.NewFastHTTPRouter()
	router.GET("/hello", fastHTTPHandler)
	fasthttp.ListenAndServe(":"+strconv.Itoa(port), router.HandleFastHTTP)
}

// go-ozzo
func ozzoHandler(c *ozzo.Context) error {
	if cpuBound {
		pow(target)
	} else {

		if sleepTime > 0 {
			time.Sleep(sleepTimeDuration)
		} else {
			runtime.Gosched()
		}
	}
	c.Write(message)

	return nil
}

func startGoozzo() {
	r := ozzo.New()
	r.Get("/hello", ozzoHandler)
	http.ListenAndServe(":"+strconv.Itoa(port), r)
}

//gowww
func startGowww() {
	rt := gowwwrouter.New()
	rt.Handle("GET", "/hello", http.HandlerFunc(helloHandler))
	http.ListenAndServe(":"+strconv.Itoa(port), rt)

}

// Gramework
func grameworkHandler(ctx *gramework.Context) {
	if cpuBound {
		pow(target)
	} else {

		if sleepTime > 0 {
			time.Sleep(sleepTimeDuration)
		} else {
			runtime.Gosched()
		}
	}

	ctx.WriteString(messageStr)
}

func startGramework() {
	gramework.SetEnv(gramework.PROD) // equivalent of ENV=prod
	app := gramework.New()
	app.GET("/hello", grameworkHandler)
	app.ListenAndServe(":" + strconv.Itoa(port))
}

// httprouter
func httpRouterHandler(w http.ResponseWriter, _ *http.Request, ps httprouter.Params) {
	if cpuBound {
		pow(target)
	} else {

		if sleepTime > 0 {
			time.Sleep(sleepTimeDuration)
		} else {
			runtime.Gosched()
		}
	}
	w.Write(message)
}
func startHTTPRouter() {
	mux := httprouter.New()
	mux.GET("/hello", httpRouterHandler)
	http.ListenAndServe(":"+strconv.Itoa(port), mux)
}

// httpTreeMux
func httpTreeMuxHandler(w http.ResponseWriter, _ *http.Request, vars map[string]string) {
	if cpuBound {
		pow(target)
	} else {

		if sleepTime > 0 {
			time.Sleep(sleepTimeDuration)
		} else {
			runtime.Gosched()
		}
	}
	w.Write(message)
}
func starthttpTreeMux() {
	mux := httptreemux.New()
	mux.GET("/hello", httpTreeMuxHandler)
	http.ListenAndServe(":"+strconv.Itoa(port), mux)
}

// lars
func larsHandler(c lars.Context) {
	if cpuBound {
		pow(target)
	} else {

		if sleepTime > 0 {
			time.Sleep(sleepTimeDuration)
		} else {
			runtime.Gosched()
		}
	}
	c.Response().Write(message)
}
func startLars() {
	mux := lars.New()
	mux.Get("/hello", larsHandler)
	http.ListenAndServe(":"+strconv.Itoa(port), mux.Serve())
}

// lion
func lionHandler(c gcontext.Context, w http.ResponseWriter, r *http.Request) {
	if cpuBound {
		pow(target)
	} else {

		if sleepTime > 0 {
			time.Sleep(sleepTimeDuration)
		} else {
			runtime.Gosched()
		}
	}
	w.Write(message)
}
func startLion() {
	mux := lion.New()
	mux.GetFunc("/hello", lionHandler)
	mux.Run(":" + strconv.Itoa(port))
}

// Macaron
func macaronHandler(c *macaron.Context) string {
	if cpuBound {
		pow(target)
	} else {

		if sleepTime > 0 {
			time.Sleep(sleepTimeDuration)
		} else {
			runtime.Gosched()
		}
	}
	return messageStr
}
func startMacaron() {
	mux := macaron.New()
	mux.Get("/hello", macaronHandler)
	http.ListenAndServe(":"+strconv.Itoa(port), mux)
}

// Martini
func martiniHandlerWrite(params martini.Params) string {
	if cpuBound {
		pow(target)
	} else {

		if sleepTime > 0 {
			time.Sleep(sleepTimeDuration)
		} else {
			runtime.Gosched()
		}
	}
	return messageStr
}

func startMartini() {
	mux := martini.NewRouter()
	mux.Get("/hello", martiniHandlerWrite)
	martini := martini.New()
	martini.Action(mux.Handle)
	http.ListenAndServe(":"+strconv.Itoa(port), martini)
}

func startMuxie() {
	mux := muxie.NewMux()
	mux.HandleFunc("/hello", helloHandler)
	http.ListenAndServe(":"+strconv.Itoa(port), mux)
}

// negroni
func startNegroni() {
	mux := http.NewServeMux()
	mux.HandleFunc("/hello", helloHandler)

	n := negroni.New()
	n.UseHandler(mux)

	http.ListenAndServe(":"+strconv.Itoa(port), n)
}

//neo
func startNeo() {
	app := neo.App()
	app.Conf.App.Addr = ":" + strconv.Itoa(port)

	app.Get("/hello", func(ctx *neo.Ctx) (int, error) {
		if cpuBound {
			pow(target)
		} else {

			if sleepTime > 0 {
				time.Sleep(sleepTimeDuration)
			} else {
				runtime.Gosched()
			}
		}
		return 200, ctx.Res.Raw(message)
	})

	app.Start()
}

// pat
func startPat() {
	mux := pat.New()
	mux.Get("/hello", http.HandlerFunc(helloHandler))
	http.ListenAndServe(":"+strconv.Itoa(port), mux)
}

//pure
func startPure() {
	p := pure.New()
	p.Get("/hello", helloHandler)
	http.ListenAndServe(":"+strconv.Itoa(port), p.Serve())
}

// R2router
func r2routerHandler(w http.ResponseWriter, req *http.Request, params r2router.Params) {
	if cpuBound {
		pow(target)
	} else {

		if sleepTime > 0 {
			time.Sleep(sleepTimeDuration)
		} else {
			runtime.Gosched()
		}
	}
	w.Write(message)
}
func startR2router() {
	mux := r2router.NewRouter()
	mux.Get("/hello", r2routerHandler)
	http.ListenAndServe(":"+strconv.Itoa(port), mux)
}

// // siris
// func sirisrouterHandler(ctx siriscontext.Context) {
// 	if sleepTime > 0 {
// 		time.Sleep(sleepTimeDuration)
// 	} else {
// 		runtime.Gosched()
// 	}
// 	ctx.HTML(messageStr)
// }
// func startSirisrouter() {
// 	app := siris.New()
// 	app.Get("/hello", sirisrouterHandler)
// 	app.Run(siris.Addr(":"+strconv.Itoa(port)), siris.WithCharset("UTF-8"))
// }

//Tango
func tangoHandler(ctx *tango.Context) {
	if cpuBound {
		pow(target)
	} else {

		if sleepTime > 0 {
			time.Sleep(sleepTimeDuration)
		} else {
			runtime.Gosched()
		}
	}
	ctx.Write(message)
}

func startTango() {
	llog.SetOutput(new(mockResponseWriter))
	llog.SetOutputLevel(llog.Lnone)

	mux := tango.NewWithLog(llog.Std)
	mux.Get("/hello", tangoHandler)
	http.ListenAndServe(":"+strconv.Itoa(port), mux)
}

//  Tiger Tonic
func startTigerTonic() {
	mux := tigertonic.NewTrieServeMux()
	mux.Handle("GET", "/hello", http.HandlerFunc(helloHandler))
	http.ListenAndServe(":"+strconv.Itoa(port), mux)
}

//  TinyRouter
func startTinyRouter() {
	routes := []tiny.Route{
		{
			Method:     "GET",
			Pattern:    "/hello",
			HandleFunc: helloHandler,
		}}
	router := tiny.New(tiny.Config{Routes: routes})
	http.ListenAndServe(":"+strconv.Itoa(port), router)
}

//  traffic
func trafficHandler(w traffic.ResponseWriter, r *traffic.Request) {
	if cpuBound {
		pow(target)
	} else {

		if sleepTime > 0 {
			time.Sleep(sleepTimeDuration)
		} else {
			runtime.Gosched()
		}
	}
	w.Write(message)
}
func startTraffic() {
	traffic.SetVar("env", "bench")
	mux := traffic.New()
	mux.Get("/hello", trafficHandler)
	http.ListenAndServe(":"+strconv.Itoa(port), mux)
}

// violetear
func startVioletear() {
	mux := violetear.New()
	mux.HandleFunc("/hello", helloHandler)
	http.ListenAndServe(":"+strconv.Itoa(port), mux)
}

// vulcan
func startVulcan() {
	mux := vulcan.NewMux()
	expr := fmt.Sprintf(`Method("%s") && Path("%s")`, "GET", "/hello")
	mux.HandleFunc(expr, helloHandler)

	http.ListenAndServe(":"+strconv.Itoa(port), mux)
}

// webgo
func getWebgoRoutes() []*webgo.Route {
	return []*webgo.Route{
		&webgo.Route{
			Name:     "hello",
			Method:   http.MethodGet,
			Pattern:  "/hello",
			Handlers: []http.HandlerFunc{helloHandler},
		},
	}
}
func startWebgo() {
	cfg := webgo.Config{
		Host:         "",
		Port:         strconv.Itoa(port),
		ReadTimeout:  120 * time.Second,
		WriteTimeout: 120 * time.Second,
	}
	router := webgo.NewRouter(&cfg, getWebgoRoutes())
	router.Start()
}

// mock
type mockResponseWriter struct{}

func (m *mockResponseWriter) Header() (h http.Header) {
	return http.Header{}
}

func (m *mockResponseWriter) Write(p []byte) (n int, err error) {
	return len(p), nil
}

func (m *mockResponseWriter) WriteString(s string) (n int, err error) {
	return len(s), nil
}

func (m *mockResponseWriter) WriteHeader(int) {}
