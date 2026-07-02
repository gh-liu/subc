package subc

import (
	"encoding/json"
	"strings"
	"testing"

	builtintemplate "github.com/gh-liu/subc/template"
)

func TestRenderBuiltinTemplateSingbox(t *testing.T) {
	nodes, err := parseShadowrocket("ss://YWVzLTI1Ni1nY206cGFzc0BleGFtcGxlLmNvbToxMjM0#SS%20Node")
	if err != nil {
		t.Fatalf("parseShadowrocket returned error: %v", err)
	}

	var out strings.Builder
	if err := builtintemplate.RenderBuiltin(&out, "singbox", nodes); err != nil {
		t.Fatalf("template RenderBuiltin returned error: %v", err)
	}
	var decoded map[string]any
	if err := json.Unmarshal([]byte(out.String()), &decoded); err != nil {
		t.Fatalf("rendered invalid json: %v\n%s", err, out.String())
	}
	outbounds := decoded["outbounds"].([]any)
	if len(outbounds) != 1 {
		t.Fatalf("got %d outbounds, want 1", len(outbounds))
	}
	ob := outbounds[0].(map[string]any)
	if ob["type"] != "shadowsocks" || ob["tag"] != "SS Node" || ob["server"] != "example.com" || ob["server_port"] != float64(1234) || ob["method"] != "aes-256-gcm" || ob["password"] != "pass" {
		t.Fatalf("unexpected outbound: %+v", ob)
	}
}
