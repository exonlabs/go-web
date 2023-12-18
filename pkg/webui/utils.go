package webui

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/exonlabs/go-web/pkg/web"
)

// //go:embed templates/* templates/macros/*
// var Tpl embed.FS

func Redirect(url string, blank bool) *web.Response {
	resp := &web.Response{
		StatusCode: http.StatusFound,
	}

	redirect := map[string]any{"redirect": url}
	if blank {
		redirect["blank"] = blank
	}

	b, err := json.Marshal(redirect)
	if err != nil {
		log.Println(err)
	}

	resp.SetHeader("Content-Type", "application/json")
	resp.Content = b
	return resp
}

func Reply(r web.Request, content string, doctitle string, params any) *web.Response {
	resp := web.NewResponse(string(content))
	if !r.IsJson() {
		return resp
	}

	jsonMap := make(map[string]any)
	if string(resp.Content) != "" {
		jsonMap["payload"] = string(resp.Content)
	}
	if doctitle != "" {
		jsonMap["doctitle"] = doctitle
	}

	var err error
	var b []byte
	if len(jsonMap) > 0 {
		b, err = json.Marshal(jsonMap)
		if err != nil {
			log.Println(err)
		}
	}

	if params != nil {
		b, err = json.Marshal(params)
		if err != nil {
			log.Println(err)
		}
	}

	resp.SetHeader("Content-Type", "application/json")
	resp.Content = b
	return resp
}

// func Notify(w web.Response, r web.Request, massage, category string,
// 	unique, sticky bool, params any) string {
// 	if params == nil {
// 		newParams := map[string]any{
// 			"notifications": []any{
// 				[]any{category, massage, unique, sticky},
// 			},
// 		}

// 		return Reply(w, r, "", newParams)
// 	}

// 	return Reply(w, r, "", nil)
// }

func RandInt(index int) string {
	min := 1
	max := index
	randVal := rand.Intn(max-min) + min
	return strconv.Itoa(randVal)
}
