package main

//go:generate statik -src=./vue/dist -dest=./ -f
import (
	"net/http"
	"strings"

	"github.com/stack-labs/stack-rpc/api"
	"github.com/stack-labs/stack-rpc/util/log"
	"github.com/stack-labs/starter-kit/pkg/plugin/opentracing"

	"github.com/stack-labs/starter-kit/console/web/beego"
	"github.com/stack-labs/starter-kit/console/web/echo"
	"github.com/stack-labs/starter-kit/console/web/gin"
	"github.com/stack-labs/starter-kit/console/web/iris"
	"github.com/stack-labs/starter-kit/console/web/statik"
	"github.com/stack-labs/starter-kit/pkg/gateway/web"
	"github.com/stack-labs/starter-kit/pkg/tracer"
)

func main() {
	md := make(map[string]string)
	md["chain"] = "gray"

	// create new web service
	service := web.NewService(
		web.Name("stack.rpc.api.console.web"),
		web.Version("latest"),
		web.Metadata(md),
	)

	// initialise service
	if err := service.Init(); err != nil {
		log.Fatal(err)
	}

	// 链路追踪
	t, closer, err := tracer.NewJaegerTracer("stack.rpc.api.console.web", "127.0.0.1:6831")
	if err != nil {
		log.Fatalf("opentracing tracer create error:%v", err)
	}
	defer closer.Close()

	// Tracing仅由Gateway控制，在下游服务中仅在有Tracing时启动
	traceHandler := opentracing.Handler(
		opentracing.WithTracer(t),
		opentracing.WithAutoStart(false),
		opentracing.WithSkipperFunc(func(r *http.Request) bool {
			return false
		}),
	)

	mux := http.NewServeMux()

	// TODO Path末尾"/"问题
	// Echo
	echoHandler, err := echo.New()
	if err != nil {
		log.Fatal(err)
	}
	mux.Handle("/v1/echo/", traceHandler(echoHandler))

	// Gin
	ginHandler, err := gin.New()
	if err != nil {
		log.Fatal(err)
	}
	mux.Handle("/v1/gin/", traceHandler(ginHandler))

	// Iris
	irisHandler, err := iris.New()
	if err != nil {
		log.Fatal(err)
	}
	mux.Handle("/v1/iris/", traceHandler(irisHandler))

	// Beego
	beegoHandler, err := beego.New()
	if err != nil {
		log.Fatal(err)
	}
	mux.Handle("/v1/beego/", traceHandler(beegoHandler))

	// register static file handler
	// 使用statik打包需要:make statik，编译时增加`-tags "statik"`标签
	mux.Handle("/", statik.Handler())

	// Path前缀
	h := &handler{
		prefix: "/console",
		mux:    mux,
	}
	service.Handle("/console/", h, &api.Endpoint{
		Name:    "console",
		Path:    []string{"^/console/*"},
		Method:  []string{"POST", "GET", "DELETE", "HEAD", "OPTIONS"},
		Handler: "proxy",
	})

	// run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}

type handler struct {
	prefix string
	mux    *http.ServeMux
}

func (h *handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Debugf("console receive request path: %v", r.URL.Path)
	r.URL.Path = strings.TrimPrefix(r.URL.Path, h.prefix)
	h.mux.ServeHTTP(w, r)
}
