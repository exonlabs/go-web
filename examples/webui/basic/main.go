package main

import (
	"flag"
	"os"
	"path/filepath"
	"runtime/debug"

	"github.com/exonlabs/go-utils/pkg/xlog"
	"github.com/exonlabs/go-web/pkg/web"
	"github.com/exonlabs/go-web/pkg/webui"
)

var (
	HOST        = "0.0.0.0"
	PORT        = 8080
	STATIC_PATH = filepath.Join("static")
	menuBuffer  = make(map[int]webui.MenuLink)
)

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
	srv.StaticPath = STATIC_PATH
	srv.SessionFactory = web.NewSessionCookieFactory(nil)

	views := []web.View{
		&IndexView{}, &HomeView{}, &NotifyView{},
		&AlertsView{}, &InputForm{}, &DatagridView{},
		&QueryBuilderView{}, &LoaderView{}, &LoginView{},
	}
	for _, view := range views {
		srv.AddView(view)
	}

	// setup webui paths
	go func() {
		err := webui.PathsSetup(STATIC_PATH)
		if err != nil {
			logger.Error("%s", err.Error())
		}
	}()

	if err := srv.Start(HOST, PORT); err != nil {
		logger.Fatal(err.Error())
		os.Exit(1)
	}
	logger.Info("exit")
}
