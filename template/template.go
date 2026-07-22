package template

import (
	"embed"
	"encoding/json"
	"fmt"
	"io"
	"io/fs"
	"strings"
	"text/template"
)

//go:embed *.gotmpl
var templates embed.FS

const suffix = ".gotmpl"

func Names() []string {
	paths, err := fs.Glob(templates, "*"+suffix)
	if err != nil {
		panic(err)
	}
	names := make([]string, len(paths))
	for i, path := range paths {
		names[i] = strings.TrimSuffix(path, suffix)
	}
	return names
}

func Render(w io.Writer, source string, data any) error {
	tpl, err := template.New("nodes").Funcs(funcs()).Parse(source)
	if err != nil {
		return err
	}
	return tpl.Execute(w, data)
}

func RenderBuiltin(w io.Writer, name string, data any) error {
	found := false
	for _, builtin := range Names() {
		if builtin == name {
			found = true
			break
		}
	}
	if !found {
		return fmt.Errorf("unknown built-in template %q", name)
	}
	source, err := templates.ReadFile(name + suffix)
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
