package subc

import (
	"io"

	builtintemplate "github.com/gh-liu/subc/template"
)

// RenderNodesTemplate renders nodes with a Go text/template source.
func RenderNodesTemplate(w io.Writer, templateSource string, nodes []Node) error {
	return builtintemplate.Render(w, templateSource, nodes)
}

// RenderBuiltinTemplate renders nodes with a named built-in template.
func RenderBuiltinTemplate(w io.Writer, name string, nodes []Node) error {
	return builtintemplate.RenderBuiltin(w, name, nodes)
}
