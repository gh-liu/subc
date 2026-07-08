package subc

import (
	"encoding/json"
	"strings"
	"testing"

	subclient "github.com/gh-liu/subc/client"
	builtintemplate "github.com/gh-liu/subc/template"
)

func TestRenderBuiltinTemplateSingbox(t *testing.T) {
	nodes, err := subclient.ParseShadowrocket("ss://YWVzLTI1Ni1nY206cGFzc0BleGFtcGxlLmNvbToxMjM0#SS%20Node")
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

func TestRenderBuiltinTemplateSingboxVLESSReality(t *testing.T) {
	nodes, err := subclient.ParseShadowrocket("vless://YXV0bzo2ZjBmNjdkYi1kZjdlLTQ3ZDktYjQ0Ny1hMjMyZTQ0MzNiNTNAaGszLm1peWF6b25vLWthb3JpLmNvbTo0NDM=?tfo=1&remark=%F0%9F%87%AD%F0%9F%87%B0Hong%20Kong%2003&tls=1&xtls=2&sni=hk4e-launcher-static.hoyoverse.com&pbk=zV4Sja0iajpdzkR1iBMbI7RRWL2nlpgwGOtirQSG-zc&sid=b81edf7fdee4&fp=chrome")
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
	ob := outbounds[0].(map[string]any)
	tls := ob["tls"].(map[string]any)
	reality := tls["reality"].(map[string]any)
	if reality["enabled"] != true || reality["public_key"] != "zV4Sja0iajpdzkR1iBMbI7RRWL2nlpgwGOtirQSG-zc" || reality["short_id"] != "b81edf7fdee4" {
		t.Fatalf("unexpected reality settings: %+v", reality)
	}
}
