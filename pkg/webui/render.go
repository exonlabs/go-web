package webui

import (
	"bytes"
	"io/fs"
	"strings"
	"text/template"
)

func Render(data any, filenames ...string) (string, error) {
	var err error
	t := tmpl(getFirstFname(filenames))
	t, err = t.ParseFiles(filenames...)
	if err != nil {
		return "", err
	}

	return executeTpl(t, data)
}

func RenderFS(fs fs.FS, data any, patterns ...string) (string, error) {
	var err error
	t := tmpl(getFirstFname(patterns))
	t, err = t.ParseFS(fs, patterns...)
	if err != nil {
		return "", err
	}

	return executeTpl(t, data)
}

func tmpl(name string) *template.Template {
	return template.New(name).
		Option("missingkey=default").Funcs(tplFunc)
}

func getFirstFname(files []string) string {
	firstFile := files[0]
	sFile := strings.Split(firstFile, "/")
	return sFile[len(sFile)-1]
}

func executeTpl(t *template.Template, data any) (string, error) {
	var html bytes.Buffer
	if err := t.Execute(&html, data); err != nil {
		return "", err
	}
	return html.String(), nil
}
