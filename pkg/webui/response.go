package webui

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/exonlabs/go-web/pkg/web"
)

// //go:embed templates/* templates/macros/*
// var Tpl embed.FS

func Redirect(url string, blank bool) (*web.Response, error) {
	resp := &web.Response{
		StatusCode: http.StatusFound,
	}

	redirect := map[string]any{"redirect": url}
	if blank {
		redirect["blank"] = blank
	}

	b, err := json.Marshal(redirect)
	if err != nil {
		return nil, err
	}

	resp.SetHeader("Content-Type", "application/json")
	resp.Content = b
	return resp, nil
}

func Reply(ctx *web.Context, content, doctitle string,
	params any) (*web.Response, error) {
	resp := web.NewResponse(string(content))
	if !ctx.Request.IsJson() {
		return resp, nil
	}

	jsonMap := make(map[string]any)
	if string(resp.Content) != "" {
		jsonMap["payload"] = string(resp.Content)
	}
	if doctitle != "" {
		jsonMap["doctitle"] = doctitle
	}

	notifications := GetFlashedMsg(ctx)
	if len(notifications) > 0 {
		var notify []any
		for _, notification := range notifications {
			for cat, msg := range notification {
				if strings.Contains(cat, ".") {
					splitedCat := strings.Split(cat, ".")
					_cat, opts := splitedCat[0], splitedCat[1]
					notify = append(notify,
						[]any{_cat, msg,
							strings.Contains(opts, "u"),
							strings.Contains(opts, "s")})
				} else {
					notify = append(notify,
						[]any{cat, msg, false, false})
				}
			}
		}
		jsonMap["notifications"] = notify
	}

	var err error
	var b []byte
	if len(jsonMap) > 0 {
		b, err = json.Marshal(jsonMap)
		if err != nil {
			return nil, err
		}
	}

	if params != nil {
		b, err = json.Marshal(params)
		if err != nil {
			return nil, err
		}
	}

	resp.SetHeader("Content-Type", "application/json")
	resp.Content = b
	return resp, nil
}

func Notify(ctx *web.Context, massage, category string, unique,
	sticky bool, params any) (*web.Response, error) {
	if params == nil {
		newParams := map[string]any{
			"notifications": []any{
				[]any{category, massage, unique, sticky},
			},
		}
		reply, err := Reply(ctx, "", "", newParams)
		if err != nil {
			return nil, err
		}
		return reply, nil
	}

	reply, err := Reply(ctx, "", "", nil)
	if err != nil {
		return nil, err
	}
	return reply, nil
}

func Flash(ctx *web.Context, category, message string) {
	flashes := []map[string]any{}

	if v, ok := ctx.Session.Get("_flashes"); ok {
		flashes = append(flashes, v.([]map[string]any)...)
	}

	flash := map[string]any{category: message}
	flashes = append(flashes, flash)
	ctx.Session.Set("_flashes", flashes)
}

func GetFlashedMsg(ctx *web.Context) []map[string]any {
	flashes := []map[string]any{}
	if v, ok := ctx.Session.Get("_flashes"); ok {
		flashes = v.([]map[string]any)
	}
	ctx.Session.Del("_flashes")
	return flashes
}
