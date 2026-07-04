package subc

import (
	"strings"
	"testing"

	subclient "github.com/gh-liu/subc/client"
	builtintemplate "github.com/gh-liu/subc/template"
)

func TestRenderNodesTemplate(t *testing.T) {
	nodes, err := subclient.ParseShadowrocket("ss://YWVzLTI1Ni1nY206cGFzc0BleGFtcGxlLmNvbToxMjM0#SS%20Node")
	if err != nil {
		t.Fatalf("parseShadowrocket returned error: %v", err)
	}

	var out strings.Builder
	err = builtintemplate.Render(&out, "{{range .}}{{.Protocol}} {{.NodeName}}\n{{end}}", nodes)
	if err != nil {
		t.Fatalf("template Render returned error: %v", err)
	}
	if out.String() != "ss SS Node\n" {
		t.Fatalf("rendered %q", out.String())
	}
}
