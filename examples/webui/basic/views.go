package main

import (
	"sync"

	"github.com/exonlabs/go-web/pkg/web"
	"github.com/exonlabs/go-web/pkg/webui"
	"github.com/exonlabs/go-web/pkg/webui/macros"
)

const cdnURL = "/static/vendor"

var (
	sharedBuff sync.Map
	count      int
)

type IndexView struct{}

func (v *IndexView) Meta() *web.ViewMeta {
	webui.AddMenulink(menuBuffer, 1, "UI Components",
		"fa-cubes", "", 0)
	return web.NewViewMeta("WebUI", "/", "/index", "/index/")
}

func (v *IndexView) DoGet(ctx *web.Context) *web.Response {
	params := map[string]any{
		"cdn_url": cdnURL,
		"langs": map[string]string{
			"en": "English",
			"ar": "العربية",
			"fr": "Français",
		},
		"doc_title": "WebUI",
		"menu":      menuBuffer,
	}

	render, err := webui.Render(params, "templates/html.tpl",
		"templates/menuboard.tpl", "templates/mainpage.tpl")
	if err != nil {
		ctx.Logger.Error(err.Error())
	}

	return web.NewResponse(render)
}

type HomeView struct{}

func (v *HomeView) Meta() *web.ViewMeta {
	webui.AddMenulink(menuBuffer, 0, "Home",
		"fa-home", "#home", 0)
	return web.NewViewMeta("Home", "/home", "/home/")
}

func (v *HomeView) DoGet(ctx *web.Context) *web.Response {
	params := map[string]any{
		"langs": map[string]string{
			"en": "English",
			"ar": "العربية",
			"fr": "Français",
		},
		"message": "Welcome",
	}

	html, err := webui.Render(params,
		"templates/option_panel.tpl")
	if err != nil {
		ctx.Logger.Error(err.Error())
	}

	return webui.Reply(*ctx.Request, html, "Home", nil)
}

type NotifyView struct{}

func (v *NotifyView) Meta() *web.ViewMeta {
	webui.AddMenulink(menuBuffer, 1, "Notifications",
		"", "#notify", 1)
	return web.NewViewMeta("Notifications", "/notify", "/notify/")
}

func (v *NotifyView) DoGet(ctx *web.Context) *web.Response {
	html, err := macros.UiAlert("general message", "showing notifications",
		true, false, "p-3")
	if err != nil {
		ctx.Logger.Error(err.Error())
	}
	html += "<script>WebUI.board_menu.show_submenu(1)</script>"

	return webui.Reply(*ctx.Request, html, "Notifications", nil)
}

type AlertsView struct{}

func (v *AlertsView) Meta() *web.ViewMeta {
	webui.AddMenulink(menuBuffer, 2, "Alerts",
		"", "#alerts", 1)
	return web.NewViewMeta("Alerts", "/alerts", "/alerts/")
}

func (v *AlertsView) DoGet(ctx *web.Context) *web.Response {
	data, err := macros.UiAlert("info", "info message",
		true, false, "px-3 pt-3")
	if err != nil {
		ctx.Logger.Error(err.Error())
	}
	html := data
	data, err = macros.UiAlert("warn", "warning message",
		true, false, "px-3")
	if err != nil {
		ctx.Logger.Error(err.Error())
	}
	html += data
	data, err = macros.UiAlert("error", "error message",
		true, false, "px-3")
	if err != nil {
		ctx.Logger.Error(err.Error())
	}
	html += data
	data, err = macros.UiAlert("success", "success message",
		true, false, "px-3")
	if err != nil {
		ctx.Logger.Error(err.Error())
	}
	html += data
	data, err = macros.UiAlert("general", "general message",
		true, false, "px-3")
	if err != nil {
		ctx.Logger.Error(err.Error())
	}
	html += data
	html += "<script>WebUI.board_menu.show_submenu(1)</script>"

	return webui.Reply(*ctx.Request, html, "Alerts", nil)
}
