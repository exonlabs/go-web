package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"runtime/debug"
	"strconv"

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
	content := "All Cookies:\n"
	for _, c := range ctx.Request.Cookies() {
		content += fmt.Sprintf("%s: %v\n", c.Name, c.Value)
	}
	return web.NewResponse(content)
}

type CookieAddView struct{}

func (v *CookieAddView) Meta() *web.ViewMeta {
	return web.NewViewMeta("add", "/add")
}

func (v *CookieAddView) DoGet(ctx *web.Context) *web.Response {
	oldVal := 0
	if c, _ := ctx.Request.Cookie("counter"); c != nil {
		oldVal, _ = strconv.Atoi(c.Value)
	}
	resp := web.RedirectResponse("/")
	resp.SetCookie(&http.Cookie{
		Name:  "counter",
		Value: strconv.Itoa(oldVal + 1),
	})
	return resp
}

type CookieSubView struct{}

func (v *CookieSubView) Meta() *web.ViewMeta {
	return web.NewViewMeta("sub", "/sub")
}

func (v *CookieSubView) DoGet(ctx *web.Context) *web.Response {
	oldVal := 0
	if c, _ := ctx.Request.Cookie("counter"); c != nil {
		oldVal, _ = strconv.Atoi(c.Value)
	}
	resp := web.RedirectResponse("/")
	resp.SetCookie(&http.Cookie{
		Name:  "counter",
		Value: strconv.Itoa(oldVal - 1),
	})
	return resp
}

type CookieDelView struct{}

func (v *CookieDelView) Meta() *web.ViewMeta {
	return web.NewViewMeta("del", "/del")
}

func (v *CookieDelView) DoGet(ctx *web.Context) *web.Response {
	resp := web.RedirectResponse("/")
	resp.DelCookie("counter")
	return resp
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
	srv.DefaultContentType = "text/plain; charset=utf-8"

	srv.AddView(&IndexView{})
	srv.AddView(&CookieAddView{})
	srv.AddView(&CookieSubView{})
	srv.AddView(&CookieDelView{})
	srv.AddView(&ExitView{})

	if err := srv.Start(HOST, PORT); err != nil {
		logger.Fatal(err.Error())
		os.Exit(1)
	}
	logger.Info("exit")
}
