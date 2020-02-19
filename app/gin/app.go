package tygin

import (
	"fmt"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	// "net/http/pprof"
	"github.com/DeanThompson/ginpprof"
	"gocommon/utils"
)

type WebApplication struct {
	engine      *gin.Engine
	debug       int64
	port        int64
	enableTLS   bool
	tlsCertFile string // TLS 证书文件
	tlsKeyFile  string // TLS key文件
	tag         string
}

func NewApplication(debug int64) *WebApplication {
	if debug == 0 {
		gin.SetMode(gin.ReleaseMode)
	} else {
		gin.SetMode(gin.DebugMode)
	}

	// gin.DebugPrintRouteFunc = func(httpMethod, absolutePath, handlerName string, nuHandlers int) {
	// 	log.WithFields(log.Fields{
	// 		"httpMethod":   httpMethod,
	// 		"absolutePath": absolutePath,
	// 		"handlerName":  handlerName,
	// 		"nuHandlers":   nuHandlers,
	// 	}).Info("gin")
	// }

	return &WebApplication{engine: gin.New(), debug: debug, tag: tyutils.UUID()}
}

func (app *WebApplication) Tag() string {
	return app.tag
}

func (app *WebApplication) EnableTLS() bool {
	return app.enableTLS
}

func (app *WebApplication) SetEnableTLS(enableTLS bool) {
	app.enableTLS = enableTLS
}

func (app *WebApplication) Debug() int64 {
	return app.debug
}

func (app *WebApplication) Engine() *gin.Engine {
	return app.engine
}

func (app *WebApplication) SetTlsFile(tlsCertFile, tlsKeyFile string) {
	app.tlsCertFile = tlsCertFile
	app.tlsKeyFile = tlsKeyFile
}

func (app *WebApplication) Init() {
	if app.debug == 1 {
		app.enablePProf()
	}

	if app.debug == 1 {
		app.engine.Use(gin.Logger())
		// app.engine.Use(gin.LoggerWithFormatter(func(params gin.LogFormatterParams) string {
		// 	return fmt.Sprintf("%s - [%s] \"%s %s %s %d %s \"%s\" %s\"\n",
		// 		params.ClientIP,
		// 		params.TimeStamp.Format(time.RFC1123),
		// 		params.Method,
		// 		params.Path,
		// 		params.Request.Proto,
		// 		params.StatusCode,
		// 		params.Latency,
		// 		params.Request.UserAgent(),
		// 		params.ErrorMessage,
		// 	)
		// }))
	}

	app.engine.Use(gin.Recovery())
}

func (app *WebApplication) Run(port int64, fn func(app *WebApplication)) error {
	app.port = port

	app.Init()

	app.engine.Use(gin.Recovery())

	fn(app)

	if app.enableTLS {
		app.runTLS(port)
	} else {
		app.run(port)
	}

	return nil
}

func (app *WebApplication) run(port int64) {
	var err error
	err = app.engine.Run(fmt.Sprintf(":%d", port))
	if err != nil {
		log.Panic(err)
	}

	// err = autotls.Run(app.engine, "www.tuyue.com", "www.tuyue1.com")
	// if err != nil {
	// 	log.Panic(err)
	// }
}

func (app *WebApplication) runTLS(port int64) {
	var err error
	err = app.engine.RunTLS(fmt.Sprintf(":%d", port), app.tlsCertFile, app.tlsKeyFile)
	if err != nil {
		log.Panic(err)
	}
}

func (app *WebApplication) enablePProf() {
	// app.engine.GET("/debug/pprof/", func(c *gin.Context) {
	// 	pprof.Index(c.Writer, c.Request)
	// })
	//
	// app.engine.GET("/debug/pprof/cmdline", func(c *gin.Context) {
	// 	pprof.Cmdline(c.Writer, c.Request)
	// })
	//
	// app.engine.GET("/debug/pprof/profile", func(c *gin.Context) {
	// 	pprof.Profile(c.Writer, c.Request)
	// })
	//
	// app.engine.GET("/debug/pprof/symbol", func(c *gin.Context) {
	// 	pprof.Symbol(c.Writer, c.Request)
	// })
	//
	// app.engine.GET("/debug/pprof/trace", func(c *gin.Context) {
	// 	pprof.Trace(c.Writer, c.Request)
	// })
	ginpprof.Wrap(app.engine)
}
