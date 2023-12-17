package main

import (
	"flag"
	"os"
	"runtime/debug"

	"github.com/exonlabs/go-utils/pkg/xlog"
	"github.com/exonlabs/go-web/pkg/web"
)

var (
	HOST = "0.0.0.0"
	PORT = 8000
)

type IndexView struct{}

func (v *IndexView) Meta() *web.ViewMeta {
	return web.NewViewMeta("index", "/")
}

func (v *IndexView) DoGet(ctx *web.Context) *web.Response {
	ctx.Logger.Info("log msg from index GET")
	return web.NewResponse("host: " + ctx.Request.Host)
}

type HomeView struct{}

func (v *HomeView) Meta() *web.ViewMeta {
	return web.NewViewMeta("home", "/home", "/home/")
}

func (v *HomeView) DoGet(ctx *web.Context) *web.Response {
	ctx.Logger.Info("log msg from home GET")
	return web.NewResponse("this is home page")
}

type RedirectView struct{}

func (v *RedirectView) Meta() *web.ViewMeta {
	return web.NewViewMeta("redirect", "/redirect")
}

func (v *RedirectView) DoGet(ctx *web.Context) *web.Response {
	ctx.Logger.Info("log msg from redirect GET")
	return web.RedirectResponse("/home")
}

type ExitView struct{}

func (v *ExitView) Meta() *web.ViewMeta {
	return web.NewViewMeta("exit", "/exit")
}

func (v *ExitView) DoGet(ctx *web.Context) *web.Response {
	ctx.Server.Stop()
	return web.NewResponse("")
}

func main() {
	logger := xlog.GetLogger()
	logger.Name = "main"
	logger.SetFormatter(xlog.NewStdFormatter(
		"{time} {level} [{source}] {message}",
		"2006-01-02 15:04:05.000000"))

	defer func() {
		if r := recover(); r != nil {
			logger.Panic("%s", r)
			logger.Trace1("%s", debug.Stack())
			logger.Warn("exit ... due to last error")
		} else {
			logger.Info("exit")
		}
	}()

	debugOpt := flag.Int("x", 0, "set debug modes, (default: 0)")
	flag.Parse()

	if *debugOpt > 0 {
		switch *debugOpt {
		case 1:
			logger.Level = xlog.DEBUG
		case 2:
			logger.Level = xlog.TRACE1
		default:
			logger.Level = xlog.TRACE2
		}
	}

	logger.Info("***** starting *****")

	srv := web.NewServer("WebPortal", nil)

	srv.AddView(&IndexView{})
	srv.AddView(&HomeView{})
	srv.AddView(&RedirectView{})
	srv.AddView(&ExitView{})

	if err := srv.Start(HOST, PORT); err != nil {
		logger.Fatal(err.Error())
		os.Exit(1)
	}
	logger.Info("exit")
}
