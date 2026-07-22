package template

import (
	"embed"
	"encoding/json"
	"fmt"
	"io"
	"text/template"
)

//go:embed *.gotmpl
var templates embed.FS

var files = map[string]string{
	"mihomo":  "mihomo.gotmpl",
	"singbox": "singbox.gotmpl",
}

func Render(w io.Writer, source string, data any) error {
	tpl, err := template.New("nodes").Funcs(funcs()).Parse(source)
	if err != nil {
		return err
	}
	return tpl.Execute(w, data)
}

func RenderBuiltin(w io.Writer, name string, data any) error {
	path, ok := files[name]
	if !ok {
		return fmt.Errorf("unknown built-in template %q", name)
	}
	source, err := templates.ReadFile(path)
	if err != nil {
		return err
	}
	return Render(w, string(source), data)
}

func funcs() template.FuncMap {
	return template.FuncMap{
		"json": func(v any) (string, error) {
			b, err := json.MarshalIndent(v, "        ", "  ")
			if err != nil {
				return "", err
			}
			return string(b), nil
		},
	}
}
