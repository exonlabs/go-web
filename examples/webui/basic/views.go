package main

import (
	"fmt"
	"math/rand"
	"slices"
	"strconv"
	"strings"
	"sync"
	"time"

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
			"fr": "Français",
			"ar": "العربية",
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
			"fr": "Français",
			"ar": "العربية",
		},
		"message": "Welcome",
	}

	html, err := webui.Render(params,
		"templates/option_panel.tpl")
	if err != nil {
		ctx.Logger.Error(err.Error())
	}

	reply, err := webui.Reply(ctx, html, "Home", nil)
	if err != nil {
		ctx.Logger.Error(err.Error())
	}

	return reply
}

type NotifyView struct{}

func (v *NotifyView) Meta() *web.ViewMeta {
	webui.AddMenulink(menuBuffer, 1, "Notifications",
		"", "#notify", 1)
	return web.NewViewMeta("Notifications", "/notify", "/notify/")
}

func (v *NotifyView) DoGet(ctx *web.Context) *web.Response {
	webui.Flash(ctx, "error.us", "error message STICKY_MSG")
	webui.Flash(ctx, "warn", "warning message")
	webui.Flash(ctx, "info", "info message")
	webui.Flash(ctx, "success", "success message")

	html, err := macros.UiAlert("general message", "showing notifications",
		true, false, "p-3")
	if err != nil {
		ctx.Logger.Error(err.Error())
	}
	html += "<script>WebUI.board_menu.show_submenu(1)</script>"

	reply, err := webui.Reply(ctx, html, "Notifications", nil)
	if err != nil {
		ctx.Logger.Error(err.Error())
	}

	return reply
}

type AlertsView struct{}

func (v *AlertsView) Meta() *web.ViewMeta {
	webui.AddMenulink(menuBuffer, 2, "Alerts",
		"", "#alerts", 1)
	return web.NewViewMeta("Alerts", "/alerts", "/alerts/")
}

func (v *AlertsView) DoGet(ctx *web.Context) *web.Response {
	data, err := macros.UiAlert("info", "info message",
		true, true, "px-3 pt-3")
	if err != nil {
		ctx.Logger.Error(err.Error())
	}
	html := data
	data, err = macros.UiAlert("warn", "warning message",
		true, true, "px-3")
	if err != nil {
		ctx.Logger.Error(err.Error())
	}
	html += data
	data, err = macros.UiAlert("error", "error message",
		true, true, "px-3")
	if err != nil {
		ctx.Logger.Error(err.Error())
	}
	html += data
	data, err = macros.UiAlert("success", "success message",
		true, true, "px-3")
	if err != nil {
		ctx.Logger.Error(err.Error())
	}
	html += data
	data, err = macros.UiAlert("general", "general message",
		false, true, "px-3")
	if err != nil {
		ctx.Logger.Error(err.Error())
	}
	html += data
	html += "<script>WebUI.board_menu.show_submenu(1)</script>"

	reply, err := webui.Reply(ctx, html, "Alerts", nil)
	if err != nil {
		ctx.Logger.Error(err.Error())
	}

	return reply
}

type InputForm struct{}

func (v *InputForm) Meta() *web.ViewMeta {
	webui.AddMenulink(menuBuffer, 3, "Input Form",
		"", "#inputform", 1)
	return web.NewViewMeta("Input Form", "/inputform", "/inputform/")
}

func (v *InputForm) DoGet(ctx *web.Context) *web.Response {
	options := map[string]any{
		"cdn_url":    cdnURL,
		"form_id":    "1234",
		"submit_url": "/inputform/",
		"fields": []map[string]any{
			{
				"type": "checkbox", "label": "Validation",
				"options": []map[string]any{
					{"label": "Server side validation", "name": "validation"},
				},
			},
			{"type": "title", "label": "Group Label"},

			{"type": "text", "label": "Required Field",
				"name": "field1", "required": true,
				"help":      "* example with input append",
				"helpguide": "Extra detailed long help for fields",
				"append": []map[string]any{
					{"type": "text", "value": ".00"},
					{"type": "text", "value": "$"},
				},
			},

			{"type": "text", "label": "Optional Field",
				"name": "field2", "required": false,
				"help":      "* example with input prepend",
				"helpguide": "Extra detailed long help for fields",
				"prepend": []map[string]any{
					{"type": "icon", "value": "fa-phone"},
					{"type": "text", "value": "+00"},
				},
			},

			{"type": "text", "label": "Optional Field",
				"name": "field3", "required": false,
				"help":      "* help text for field",
				"helpguide": "Extra detailed long help for field",
				"append": []map[string]any{
					{"type": "select", "options": []map[string]any{
						{"label": ".com", "value": ".com"},
						{"label": ".net", "value": ".net", "selected": true},
						{"label": ".org", "value": ".org"}}}},
			},

			{"type": "textarea", "label": "Textarea",
				"name":      "field4",
				"help":      "* help text for field",
				"helpguide": "Extra detailed long help for fields"},

			{"type": "title", "label": "Group Label"},

			{"type": "password", "label": "Password 1",
				"name": "pass1", "required": true, "strength": true},

			{"type": "password", "label": "Password 2",
				"name": "pass2", "required": true, "confirm": true},

			{"type": "password", "label": "Password 3",
				"name": "pass3", "required": true, "strength": true,
				"confirm": true},

			{"type": "title", "label": "Group Label"},

			{"type": "select", "label": "Select",
				"name": "select1", "required": true,
				"options": []map[string]any{{"label": "Select", "value": nil},
					{"label": "Option 1", "value": "01"},
					{"label": "Option 2", "value": "02"},
					{"label": "Option 3", "value": "03"},
					{"label": "Option 4", "value": "04"},
					{"label": "Option 5", "value": "05"}}},

			{"type": "select", "label": "Select multiple",
				"name": "select2", "required": true, "multiple": true,
				"options": []map[string]any{{"label": "Option 1", "value": "01",
					"selected": true},
					{"label": "Option 2", "value": "02",
						"selected": true},
					{"label": "Option 3", "value": "03"},
					{"label": "Option 4", "value": "04"},
					{"label": "Option 5", "value": "05"}}},

			{"type": "title", "label": "Group Label"},

			{"type": "checkbox", "label": "Checkbox",
				"helpguide": "Extra detailed long help for fields",
				"options": []map[string]any{{"label": "Select 1",
					"name": "check1", "selected": true},
					{"label": "Select 2",
						"name": "check2"}}},

			{"type": "radio", "label": "Radio",
				"name": "radio1", "required": true,
				"helpguide": "Extra detailed long help for fields",
				"options": []map[string]any{{"label": "Option 1", "value": "1"},
					{"label": "Option 2", "value": "2"},
					{"label": "Option 3", "value": "3"}}},

			{"type": "title", "label": "Group Label"},

			{"type": "datetime", "label": "Date & Time",
				"name": "date1", "required": true},

			{"type": "date", "label": "Date", "name": "date2"},

			{"type": "time", "label": "Time", "name": "time1"},

			{"type": "title", "label": "Group Label"},

			{"type": "file", "label": "Upload File",
				"name": "files1", "required": true,
				"format": ".txt,.pdf,.png",
				"help":   fmt.Sprintf("* %s: <span dir=\"ltr\">.txt, .pdf, .png</span>", "allowed types")},

			{"type": "file", "label": "Upload Multiple",
				"name": "files2", "required": false, "multiple": true,
				"placeholder": "", "maxsize": 1048576,
				"help":      "* all types allowed, max file size: 1MB",
				"helpguide": "Extra detailed long help for fields"},
		},
	}

	form, err := macros.UiInputForm(options, "")

	if err != nil {
		ctx.Logger.Error(err.Error())
	}
	html, err := webui.Render(map[string]any{"contents": form},
		"templates/input_form.tpl")

	if err != nil {
		ctx.Logger.Error(err.Error())
	}

	html += "<script>WebUI.board_menu.show_submenu(1)</script>"

	reply, err := webui.Reply(ctx, html, "Input Form", nil)
	if err != nil {
		ctx.Logger.Error(err.Error())
	}

	return reply
}

func (v *InputForm) DoPost(ctx *web.Context) *web.Response {
	validation := ctx.Request.FormValue("validation")
	if validation == "1" {
		params := map[string]any{
			"validation": []string{"field1", "date1", "files1"},
		}
		reply, err := webui.Reply(ctx, "", "", params)
		if err != nil {
			ctx.Logger.Error(err.Error())
		}

		return reply
	}

	msg := fmt.Sprintf(`%s<br><div dir="ltr" style="text-align:left">`,
		"Submited Data:")
	for k := range ctx.Request.Form {
		if k == "_csrf_token" {
			continue
		}

		v := ctx.Request.FormValue(k)
		msg += fmt.Sprintf("<b>%s:</b> %s<br>", k, v)
	}

	for k := range ctx.Request.MultipartForm.File {
		_, fh, err := ctx.Request.FormFile(k)
		if err != nil {
			ctx.Logger.Error(err.Error())
		}

		msg += fmt.Sprintf("<b>%s:</b> %s<br>",
			fh.Filename, fh.Header.Get("Content-Type"))
	}

	msg += "</div>"
	msg = strings.ReplaceAll(msg, "'", "")

	notify, err := webui.Notify(ctx, msg, "success", false, true, nil)
	if err != nil {
		ctx.Logger.Error(err.Error())
	}

	return notify

}

type DatagridView struct{}

func (v *DatagridView) Meta() *web.ViewMeta {
	webui.AddMenulink(menuBuffer, 4, "Datagrid",
		"", "#datagrid", 1)
	return web.NewViewMeta("datagrid", "/datagrid", "/datagrid/")
}

func (v *DatagridView) DoGet(ctx *web.Context) *web.Response {
	options := map[string]any{
		"cdn_url":     cdnURL,
		"grid_id":     "1234",
		"base_url":    "/datagrid",
		"load_url":    "/datagrid/loaddata",
		"length_menu": []string{"10", "50", "100", "250", "-1"},
		"columns": []map[string]any{
			{"id": "field1", "title": "Field Name 1"},
			{"id": "field2", "title": "Field Name 2"},
			{"id": "field3.item1", "title": "Field3_1"},
			{"id": "field3.item2", "title": "Field3_2"},
			{"id": "field4", "title": "Other Field 4"},
			{"id": "field5", "title": "Extra Field 5",
				"visible": false},
			{"id": "field6", "title": "Data Field 6",
				"visible": false},
			{"id": "field7", "title": "Field Header 7",
				"visible": false},
			{"id": "field8", "title": "Field8",
				"visible": false},
			{"id": "field9", "title": "Field Number 9",
				"visible": false},
		},

		"export": map[string]any{
			"types":              []string{"csv", "xls", "print"},
			"file_title":         "Example Data",
			"file_prefix":        "export",
			"csv_fieldSeparator": ";",
			"csv_fieldBoundary":  "",
		},

		"single_ops": []map[string]any{
			{"label": "Single Operation 1", "action": "single_op1"},
			{"label": "Single Op 2 with confirm", "action": "single_op2",
				"confirm": "Are you sure you want to do this operation?"},
		},

		"group_ops": []map[string]any{
			{"label": "Group Operation 1", "action": "group_op1"},
			{"label": "Group Op 2 with confirm", "action": "group_op2",
				"confirm": "Are you sure?"},
			{"label": "Op 3 with Reload", "action": "group_op3"},
		},
	}

	dataGrid, err := macros.UiStdDataGrid(options, "")
	if err != nil {
		ctx.Logger.Error(err.Error())
	}

	html, err := webui.Render(map[string]any{"contents": dataGrid},
		"templates/data_grid.tpl")

	if err != nil {
		ctx.Logger.Error(err.Error())
	}

	html += "<script>WebUI.board_menu.show_submenu(1)</script>"
	reply, err := webui.Reply(ctx, html, "Datagrid", nil)
	if err != nil {
		ctx.Logger.Error(err.Error())
	}

	return reply

}

func (v *DatagridView) DoPost(ctx *web.Context) *web.Response {
	action := strings.Split(
		strings.TrimPrefix(ctx.Request.URL.Path, "/"), "/")[1]

	if action == "loaddata" {
		data := []map[string]any{}

		for k := 0; k < 228; k++ {
			_k := fmt.Sprintf("%03d", k)

			d := map[string]any{
				"DT_RowId": fmt.Sprintf("rowid_%s", _k),
				"field1": macros.StdDataGridLink(
					fmt.Sprintf("master_%s", _k), "#datagrid", "", "", nil),
				"field2": macros.StdDataGridText(
					fmt.Sprintf("field2 %s", _k), "", "", nil),
			}

			item := map[string]any{}

			if rand.Intn(3) == 0 {
				item["item1"] = macros.StdDataGridText(
					fmt.Sprintf("field3.1 %s", _k), "", "", nil)
			} else {
				item["item1"] = macros.StdDataGridText("'", "", "", nil)
			}

			if rand.Intn(4) == 0 {
				item["item2"] = macros.StdDataGridText(
					fmt.Sprintf("field3.2 %s", _k), "", "", nil)
			} else {
				item["item2"] = macros.StdDataGridText("'", "", "", nil)
			}

			d["field3"] = item

			var pill string

			if rand.Intn(2) == 0 {
				pill = "Yes"
			} else {
				pill = "No"
			}

			d["field4"] = macros.StdDataGridPill(pill, "Yes", "", nil)
			d["field5"] = macros.StdDataGridCheck(rand.Intn(2) == 0, "", nil)

			if rand.Intn(2) == 0 {
				d["field6"] = macros.StdDataGridText(
					fmt.Sprintf("field6 %s", _k), "", "", nil)
			} else {
				d["field6"] = macros.StdDataGridText("'", "", "", nil)
			}

			if rand.Intn(2) == 0 {
				d["field7"] = macros.StdDataGridText(
					fmt.Sprintf("field7 %s", _k), "", "", nil)
			} else {
				d["field7"] = macros.StdDataGridText("'", "", "", nil)
			}

			if rand.Intn(2) == 0 {
				d["field8"] = macros.StdDataGridText(
					fmt.Sprintf("field8 %s", _k), "", "", nil)
			} else {
				d["field8"] = macros.StdDataGridText("'", "", "", nil)
			}

			if rand.Intn(2) == 0 {
				d["field9"] = macros.StdDataGridText(
					fmt.Sprintf("field9 %s", _k), "", "", nil)
			} else {
				d["field9"] = macros.StdDataGridText("'", "", "", nil)
			}

			data = append(data, d)
		}

		reply, err := webui.Reply(ctx, "", "", map[string]any{"payload": data})
		if err != nil {
			ctx.Logger.Error(err.Error())
		}

		return reply
	}

	if slices.Contains([]string{"single_op1", "single_op2",
		"group_op1", "group_op2", "group_op3"}, action) {
		ctx.Request.ParseForm()
		rows := ctx.Request.Form["items[]"]

		if len(rows) > 20 {
			rows = rows[:21]
			rows[20] = "..."
		}

		msg := `<span dir="ltr" style="text-align:left">`
		msg += fmt.Sprintf("Operation: %s<br>", action)
		msg += fmt.Sprintf("Rows: %s", rows)
		msg += "</span>"
		msg = strings.ReplaceAll(msg, "'", "")
		notify, err := webui.Notify(ctx, msg, "success", false, true, nil)
		if err != nil {
			ctx.Logger.Error(err.Error())
		}
		return notify
	}

	notify, err := webui.Notify(ctx, "Invalid request", "error", false, true, nil)
	if err != nil {
		ctx.Logger.Error(err.Error())
	}

	return notify
}

type QueryBuilderView struct{}

func (v *QueryBuilderView) Meta() *web.ViewMeta {
	webui.AddMenulink(menuBuffer, 5, "Query Builder",
		"", "#qbuilder", 1)
	return web.NewViewMeta("qbuilder", "/qbuilder", "/qbuilder/")
}

func (v *QueryBuilderView) DoGet(ctx *web.Context) *web.Response {
	options := map[string]any{
		"cdn_url": cdnURL,
		"form_id": "1234",
		"filters": []map[string]any{
			{"id": "field1", "label": "Field 1", "type": "string",
				"operators": []string{"equal", "not_equal", "contains"}},
			{"id": "field2", "label": "Field Name 2", "type": "string",
				"input": "textarea"},
			{"id": "field3_1", "label": "Integer 1", "type": "integer",
				"input": "text"},
			{"id": "field3_2", "label": "Integer 2", "type": "integer",
				"input": "number"},
			{"id": "field3_3", "label": "Double 1", "type": "double",
				"input": "number"},
			{"id": "field4", "label": "Select", "type": "integer",
				"input":     "select",
				"values":    map[int]string{1: "Option 1", 2: "Option 2", 3: "Option 3"},
				"operators": []string{"equal", "not_equal", "in", "not_in"}},
			{"id": "field5", "label": "Checkbox", "type": "integer",
				"input": "radio", "values": []map[int]any{{1: "Yes"}, {0: "No"}},
				"operators": []string{"equal"}},
			{"id": "field6", "label": "Choose", "type": "integer",
				"input": "checkbox",
				"values": []map[int]any{{1: "Opt 1"}, {2: "Opt 2"}, {3: "Opt 3"},
					{4: "Opt 4"}, {5: "Opt 5"}}},
		},

		"initial_rules": map[string]any{
			"not":       true,
			"condition": "AND",
			"rules": []map[string]any{
				{"id": "field1", "operator": "equal", "value": "value"},
				{"id": "field3_2", "operator": "less", "value": 10},
				{
					"condition": "OR",
					"rules": []map[string]any{
						{"id": "field1", "operator": "equal",
							"value": "value2"},
						{"id": "field6", "operator": "equal",

							"value": 2},
					},
				},

				{
					"not":       true,
					"condition": "OR",
					"rules": []map[string]any{
						{"id": "field4", "operator": "equal",
							"value": 3},
						{"id": "field2", "operator": "not_equal",
							"value": "text value"},
					},
				},
			},
		},
	}

	queryBuilder, err := macros.UiQBuilder(options, "")

	if err != nil {
		ctx.Logger.Error(err.Error())
	}

	html, err := webui.Render(map[string]any{"contents": queryBuilder},
		"templates/query_builder.tpl")

	if err != nil {
		ctx.Logger.Error(err.Error())
	}

	html += "<script>WebUI.board_menu.show_submenu(1)</script>"
	reply, err := webui.Reply(ctx, html, "Query Builder", nil)
	if err != nil {
		ctx.Logger.Error(err.Error())
	}

	return reply
}

type LoaderView struct{}

func (v *LoaderView) Meta() *web.ViewMeta {
	webui.AddMenulink(menuBuffer, 2, "Page Loader",
		"", "#loader", 0)
	return web.NewViewMeta("loader", "/loader", "/loader/")
}

func (v *LoaderView) DoGet(ctx *web.Context) *web.Response {
	html, err := macros.UiAlert("message", "loaded after delay",
		false, false, "p-3")

	if err != nil {
		ctx.Logger.Error(err.Error())
	}

	render, err := webui.Render(nil,
		"templates/progress_loader.tpl")

	if err != nil {
		ctx.Logger.Error(err.Error())
	}

	html += render

	// simulate delay
	time.Sleep(time.Second * 1)
	reply, err := webui.Reply(ctx, html, "Page Loader", nil)
	if err != nil {
		ctx.Logger.Error(err.Error())
	}

	return reply
}

func (v *LoaderView) DoPost(ctx *web.Context) *web.Response {
	count++
	ctx.Request.ParseForm()

	// get loading progress status
	if ctx.Request.Form.Get("get_progress") != "" {
		res, _ := sharedBuff.Load("loader_progress")
		reply, err := webui.Reply(ctx, strconv.Itoa(res.(int)), "", nil)
		if err != nil {
			ctx.Logger.Error(err.Error())
		}

		return reply
	}

	// simulate long delay
	t := 5

	for i := 1; i < t; i++ {
		sharedBuff.Store("loader_progress", i*100/t)
		time.Sleep(time.Duration(float64(time.Second) * 1))
	}

	sharedBuff.Delete("loader_progress")
	count = 0
	notify, err := webui.Notify(ctx, "success", "success",
		false, false, nil)
	if err != nil {
		ctx.Logger.Error(err.Error())
	}

	return notify
}

type LoginView struct{}

func (v *LoginView) Meta() *web.ViewMeta {
	webui.AddMenulink(menuBuffer, 3, "Login Page",
		"", "loginpage", 0)
	return web.NewViewMeta("loginpage", "/loginpage/")
}

func (v *LoginView) DoGet(ctx *web.Context) *web.Response {
	action := strings.Split(
		strings.TrimPrefix(ctx.Request.URL.Path, "/"), "/")[1]

	if action == "load" {
		html, err := macros.UiLoginForm(map[string]any{
			"cdn_url":    cdnURL,
			"submit_url": "/loginpage/",
			"authkey":    "123456",
		}, "text-white bg-secondary")
		if err != nil {
			ctx.Logger.Error(err.Error())
		}

		reply, err := webui.Reply(ctx, html, "Loginpage", nil)
		if err != nil {
			ctx.Logger.Error(err.Error())
		}

		return reply
	} else {
		params := map[string]any{
			"cdn_url":   cdnURL,
			"doc_title": "WebUI",
			"langs": map[string]string{
				"en": "English",
				"fr": "Français",
				"ar": "العربية",
			},

			"load_url": "loginpage/load",
		}

		html, err := webui.Render(params,
			"templates/html.tpl",
			"templates/simplepage.tpl",
			"templates/loginpage.tpl")
		if err != nil {
			ctx.Logger.Error(err.Error())
		}

		return web.NewResponse(html)
	}
}

func (v *LoginView) DoPost(ctx *web.Context) *web.Response {
	ctx.Request.ParseForm()
	username := ctx.Request.Form.Get("username")
	authdigest := ctx.Request.Form.Get("digest")

	var errMsg string
	if username == "" || authdigest == "" {
		errMsg = "Please enter username and password"
	} else {
		if username == "admin" {
			notify, err := webui.Notify(ctx, "Welcome, admin",
				"success", false, true, nil)
			if err != nil {
				ctx.Logger.Error(err.Error())
			}

			return notify
		} else {
			errMsg = "Invalid username or password"
		}
	}

	notify, err := webui.Notify(ctx, errMsg,
		"error", false, true, nil)
	if err != nil {
		ctx.Logger.Error(err.Error())
	}

	return notify
}
