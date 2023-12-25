package webui

import (
	"strings"
	"text/template"

	"slices"
)

var tplFunc = template.FuncMap{
	"join":       strings.Join,
	"replaceAll": strings.ReplaceAll,
	"inStr":      inStrSlice,
}

func inStrSlice(s []string, v string) bool {
	return slices.Contains(s, v)
}
