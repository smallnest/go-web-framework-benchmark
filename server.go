package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"runtime"
	"strconv"
	"time"

	"github.com/Unknwon/macaron"
	"github.com/ant0ine/go-json-rest/rest"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"github.com/bmizerany/pat"
	"github.com/buaazp/fasthttprouter"
	"github.com/claygod/Bxog"
	"github.com/dimfeld/httptreemux"
	"github.com/dinever/golf"
	"github.com/emicklei/go-restful"
	"github.com/gin-gonic/gin"
	"github.com/go-gas/gas"
	"github.com/go-martini/martini"
	ozzo "github.com/go-ozzo/ozzo-routing"
	"github.com/go-playground/lars"
	"github.com/go-zoo/bone"
	"github.com/gocraft/web"
	"github.com/gorilla/mux"
	"github.com/ivpusic/neo"
	"github.com/julienschmidt/httprouter"
	echov3 "github.com/labstack/echo"
	echov3fasthttp "github.com/labstack/echo/engine/fasthttp"
	echov3standard "github.com/labstack/echo/engine/standard"
	llog "github.com/lunny/log"
	"github.com/lunny/tango"
	vulcan "github.com/mailgun/route"
	"github.com/mikespook/possum"
	possumrouter "github.com/mikespook/possum/router"
	possumview "github.com/mikespook/possum/view"
	"github.com/mustafaakin/gongular"
	"github.com/naoina/denco"
	"github.com/pilu/traffic"
	"github.com/plimble/ace"
	"github.com/pressly/chi"
	routing "github.com/qiangxue/fasthttp-routing"
	"github.com/rcrowley/go-tigertonic"
	"github.com/valyala/fasthttp"
	"github.com/vanng822/r2router"
	goji "goji.io"
	gojipat "goji.io/pat"
	gcontext "golang.org/x/net/context"
	"gopkg.in/baa.v1"
	lion "gopkg.in/celrenheit/lion.v1"
	// guavaweb "github.com/GuavaStudio/web"
)

var port = 8080
var sleepTime = 0
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
	case "ace":
		startAce()
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
	case "echov3standard":
		startEchoV3Standard()
	case "echov3fasthttp":
		startEchoV3Fasthttp()
	case "fasthttp-raw":
		startFasthttp()
	case "fasthttprouter":
		startFastHTTPRouter()
	case "fasthttp-routing":
		startFastHTTPRouting()
	case "gas":
		startGas()
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
	case "go-ozzo":
		startGoozzo()
	// case "guavastudio_web":
	// 	startGuavaStudioWeb()
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
	case "neo":
		startNeo()
	case "pat":
		startPat()
	case "possum":
		startPossum()
	case "r2router":
		startR2router()
	case "tango":
		startTango()
	case "tiger":
		startTigerTonic()
	case "traffic":
		startTraffic()
	case "vulcan":
		startVulcan()
	}
}

// default mux
func helloHandler(w http.ResponseWriter, r *http.Request) {
	if sleepTime > 0 {
		time.Sleep(sleepTimeDuration)
	}
	w.Write(message)
}
func startDefaultMux() {
	http.HandleFunc("/hello", helloHandler)
	http.ListenAndServe(":"+strconv.Itoa(port), nil)
}

//ace
func aceHandler(c *ace.C) {
	if sleepTime > 0 {
		time.Sleep(sleepTimeDuration)
	}
	c.Writer.Write(message)
}
func startAce() {
	mux := ace.New()
	mux.GET("/hello", aceHandler)
	mux.Run(":" + strconv.Itoa(port))
}

// baa
func baaHandler(ctx *baa.Context) {
	if sleepTime > 0 {
		time.Sleep(sleepTimeDuration)
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
	if sleepTime > 0 {
		time.Sleep(sleepTimeDuration)
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
 	if sleepTime > 0 {
 		time.Sleep(sleepTimeDuration)
 	}
 	io.WriteString(w, message)
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
	if sleepTime > 0 {
		time.Sleep(sleepTimeDuration)
	}
	w.Write(message)
}
func startDenco() {
	mux := denco.NewMux()
	handler, _ := mux.Build([]denco.Handler{mux.GET("/hello", denco.HandlerFunc(dencoHandler))})
	http.ListenAndServe(":"+strconv.Itoa(port), handler)
}

// echov3-standard
func echov3Handler(c echov3.Context) error {
	if sleepTime > 0 {
		time.Sleep(sleepTimeDuration)
	}
	c.Response().Write(message)
	return nil
}
func startEchoV3Standard() {
	mux := echov3.New()
	mux.Get("/hello", echov3Handler)
	mux.Run(echov3standard.New(":" + strconv.Itoa(port)))
}

// echov3-fasthttp
func startEchoV3Fasthttp() {
	mux := echov3.New()
	mux.Get("/hello", echov3Handler)
	mux.Run(echov3fasthttp.New(":" + strconv.Itoa(port)))
}

//fasthttp
func fastHTTPRawHandler(ctx *fasthttp.RequestCtx) {
	if string(ctx.Method()) == "GET" {
		switch string(ctx.Path()) {
		case "/hello":
			if sleepTime > 0 {
				time.Sleep(sleepTimeDuration)
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
	if sleepTime > 0 {
		time.Sleep(sleepTimeDuration)
	}
	ctx.Write(message)
}
func startFastHTTPRouter() {
	mux := fasthttprouter.New()
	mux.GET("/hello", fastHTTPHandler)
	fasthttp.ListenAndServe(":"+strconv.Itoa(port), mux.Handler)
}

//fasthttprouting
func fastHTTPRoutingHandler(c *routing.Context) error {
	if sleepTime > 0 {
		time.Sleep(sleepTimeDuration)
	}
	c.Write(message)
	return nil
}
func startFastHTTPRouting() {
	mux := routing.New()
	mux.Get("/hello", fastHTTPRoutingHandler)
	fasthttp.ListenAndServe(":"+strconv.Itoa(port), mux.HandleRequest)
}

//gas
func startGas() {
	g := gas.New()
	g.Router.Get("/hello", func(c *gas.Context) error {
		if sleepTime > 0 {
			time.Sleep(sleepTimeDuration)
		}
		//c.Write(message)
		return c.STRING(200, messageStr)
	})
	g.Run(":" + strconv.Itoa(port))
}

// gin
func ginHandler(c *gin.Context) {
	if sleepTime > 0 {
		time.Sleep(sleepTimeDuration)
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
	if sleepTime > 0 {
		time.Sleep(sleepTimeDuration)
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
	if sleepTime > 0 {
		time.Sleep(sleepTimeDuration)
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
	if sleepTime > 0 {
		time.Sleep(sleepTimeDuration)
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
	if sleepTime > 0 {
		time.Sleep(sleepTimeDuration)
	}
	ctx.Send(messageStr)
}

func startGolf() {
	app := golf.New()
	app.Get("/hello", golfHandler)
	app.Run(":" + strconv.Itoa(port))
}

func startGongular() {
	g := gongular.NewRouter()
	g.DisableDebug()
	g.InfoLog.SetOutput(ioutil.Discard)
	g.InfoLog.SetFlags(0)
	g.GET("/hello", func(c *gongular.Context) string {
		if sleepTime > 0 {
			time.Sleep(sleepTimeDuration)
		}
		return messageStr
	})
	g.ListenAndServe(":" + strconv.Itoa(port))
}

// goRestful
func goRestfulHandler(r *restful.Request, w *restful.Response) {
	if sleepTime > 0 {
		time.Sleep(sleepTimeDuration)
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

// go-ozzo
func ozzoHandler(c *ozzo.Context) error {
	if sleepTime > 0 {
		time.Sleep(sleepTimeDuration)
	}
	c.Write(message)

	return nil
}

func startGoozzo() {
	r := ozzo.New()
	r.Get("/hello", ozzoHandler)
	http.ListenAndServe(":"+strconv.Itoa(port), r)
}

// //GuavaStudio/Web
// func hello(val string) string {
// 	if sleepTime > 0 {
// 		time.Sleep(sleepTimeDuration)
// 	}

// 	return messageStr
// }

// func startGuavaStudioWeb() {
// 	guavaweb.Get("/(.*)", hello)
// 	guavaweb.Run(":" + strconv.Itoa(port))
// }

// httprouter
func httpRouterHandler(w http.ResponseWriter, _ *http.Request, ps httprouter.Params) {
	if sleepTime > 0 {
		time.Sleep(sleepTimeDuration)
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
	if sleepTime > 0 {
		time.Sleep(sleepTimeDuration)
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
	if sleepTime > 0 {
		time.Sleep(sleepTimeDuration)
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
	if sleepTime > 0 {
		time.Sleep(sleepTimeDuration)
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
	if sleepTime > 0 {
		time.Sleep(sleepTimeDuration)
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
	if sleepTime > 0 {
		time.Sleep(sleepTimeDuration)
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

//neo
func startNeo() {
	app := neo.App()
	app.Conf.App.Addr = ":" + strconv.Itoa(port)

	app.Get("/hello", func(ctx *neo.Ctx) (int, error) {
		if sleepTime > 0 {
			time.Sleep(sleepTimeDuration)
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

// Possum
func possumHandler(c *possum.Context) error {
	if sleepTime > 0 {
		time.Sleep(sleepTimeDuration)
	}
	c.Response.Write(message)
	return nil
}
func startPossum() {
	mux := possum.NewServerMux()
	mux.HandleFunc(possumrouter.Simple("/hello"), possumHandler, possumview.Simple("text/html", "utf-8"))
	http.ListenAndServe(":"+strconv.Itoa(port), mux)
}

// R2router
func r2routerHandler(w http.ResponseWriter, req *http.Request, params r2router.Params) {
	if sleepTime > 0 {
		time.Sleep(sleepTimeDuration)
	}
	w.Write(message)
}
func startR2router() {
	mux := r2router.NewRouter()
	mux.Get("/hello", r2routerHandler)
	http.ListenAndServe(":"+strconv.Itoa(port), mux)
}

//Tango
func tangoHandler(ctx *tango.Context) {
	if sleepTime > 0 {
		time.Sleep(sleepTimeDuration)
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

//  traffic
func trafficHandler(w traffic.ResponseWriter, r *traffic.Request) {
	if sleepTime > 0 {
		time.Sleep(sleepTimeDuration)
	}
	w.Write(message)
}
func startTraffic() {
	traffic.SetVar("env", "bench")
	mux := traffic.New()
	mux.Get("/hello", trafficHandler)
	http.ListenAndServe(":"+strconv.Itoa(port), mux)
}

// vulcan
func startVulcan() {
	mux := vulcan.NewMux()
	expr := fmt.Sprintf(`Method("%s") && Path("%s")`, "GET", "/hello")
	mux.HandleFunc(expr, helloHandler)

	http.ListenAndServe(":"+strconv.Itoa(port), mux)
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
