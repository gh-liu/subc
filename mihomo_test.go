package subc

import (
	"strings"
	"testing"

	subclient "github.com/gh-liu/subc/client"
	"github.com/gh-liu/subc/protocol"
	builtintemplate "github.com/gh-liu/subc/template"
	"gopkg.in/yaml.v3"
)

func TestRenderBuiltinTemplateMihomo(t *testing.T) {
	nodes, err := subclient.ParseShadowrocket("ss://YWVzLTI1Ni1nY206cGFzc0BleGFtcGxlLmNvbToxMjM0#SS%20Node")
	if err != nil {
		t.Fatalf("parseShadowrocket returned error: %v", err)
	}

	var out strings.Builder
	if err := builtintemplate.RenderBuiltin(&out, "mihomo", nodes); err != nil {
		t.Fatalf("template RenderBuiltin returned error: %v", err)
	}
	if !strings.HasPrefix(out.String(), "proxies:\n") {
		t.Fatalf("rendered output is not a Mihomo YAML fragment:\n%s", out.String())
	}
	var decoded map[string]any
	if err := yaml.Unmarshal([]byte(out.String()), &decoded); err != nil {
		t.Fatalf("rendered invalid YAML: %v\n%s", err, out.String())
	}
	if len(decoded) != 1 {
		t.Fatalf("rendered top-level fields %+v, want only proxies", decoded)
	}
	proxies := decoded["proxies"].([]any)
	if len(proxies) != 1 {
		t.Fatalf("got %d proxies, want 1", len(proxies))
	}
	proxy := proxies[0].(map[string]any)
	if proxy["type"] != "ss" || proxy["name"] != "SS Node" || proxy["server"] != "example.com" || proxy["port"] != 1234 || proxy["cipher"] != "aes-256-gcm" || proxy["password"] != "pass" || proxy["udp"] != true {
		t.Fatalf("unexpected proxy: %+v", proxy)
	}
}

func TestRenderBuiltinTemplateMihomoVLESSReality(t *testing.T) {
	nodes, err := subclient.ParseShadowrocket("vless://YXV0bzo2ZjBmNjdkYi1kZjdlLTQ3ZDktYjQ0Ny1hMjMyZTQ0MzNiNTNAaGszLm1peWF6b25vLWthb3JpLmNvbTo0NDM=?tfo=1&remark=%F0%9F%87%AD%F0%9F%87%B0Hong%20Kong%2003&tls=1&xtls=2&sni=hk4e-launcher-static.hoyoverse.com&pbk=zV4Sja0iajpdzkR1iBMbI7RRWL2nlpgwGOtirQSG-zc&sid=b81edf7fdee4&fp=chrome")
	if err != nil {
		t.Fatalf("parseShadowrocket returned error: %v", err)
	}

	var out strings.Builder
	if err := builtintemplate.RenderBuiltin(&out, "mihomo", nodes); err != nil {
		t.Fatalf("template RenderBuiltin returned error: %v", err)
	}
	var decoded map[string]any
	if err := yaml.Unmarshal([]byte(out.String()), &decoded); err != nil {
		t.Fatalf("rendered invalid YAML: %v\n%s", err, out.String())
	}
	proxy := decoded["proxies"].([]any)[0].(map[string]any)
	if proxy["type"] != "vless" || proxy["flow"] != "xtls-rprx-vision" || proxy["tls"] != true || proxy["servername"] != "hk4e-launcher-static.hoyoverse.com" || proxy["client-fingerprint"] != "chrome" {
		t.Fatalf("unexpected VLESS proxy: %+v", proxy)
	}
	reality := proxy["reality-opts"].(map[string]any)
	if reality["public-key"] != "zV4Sja0iajpdzkR1iBMbI7RRWL2nlpgwGOtirQSG-zc" || reality["short-id"] != "b81edf7fdee4" {
		t.Fatalf("unexpected reality options: %+v", reality)
	}
}

func TestRenderBuiltinTemplateMihomoEmpty(t *testing.T) {
	var out strings.Builder
	if err := builtintemplate.RenderBuiltin(&out, "mihomo", nil); err != nil {
		t.Fatalf("template RenderBuiltin returned error: %v", err)
	}
	var decoded map[string]any
	if err := yaml.Unmarshal([]byte(out.String()), &decoded); err != nil {
		t.Fatalf("rendered invalid YAML: %v\n%s", err, out.String())
	}
	if proxies := decoded["proxies"].([]any); len(proxies) != 0 {
		t.Fatalf("got %d proxies, want 0", len(proxies))
	}
}

func TestRenderBuiltinTemplateMihomoSupportedProtocols(t *testing.T) {
	nodes := []protocol.Node{
		protocol.SSRNode{Base: protocol.Base{Type: "ssr", Name: "SSR"}, Host: "ssr.example.com", Port: 443, Method: "aes-256-cfb", Password: "pass", Params: map[string]string{"protocol": "auth_sha1_v4", "obfs": "tls1.2_ticket_auth"}},
		protocol.VMessNode{Base: protocol.Base{Type: "vmess", Name: "VMess"}, Host: "vmess.example.com", Port: 443, UUID: "uuid", Params: map[string]string{"aid": "0", "net": "ws", "path": "/ws", "host": "cdn.example.com", "tls": "tls"}},
		protocol.TrojanNode{Base: protocol.Base{Type: "trojan", Name: "Trojan"}, Host: "trojan.example.com", Port: 443, Password: "pass", Params: map[string]string{"sni": "example.com"}},
		protocol.Hysteria2Node{Base: protocol.Base{Type: "hysteria2", Name: "Hysteria2"}, Host: "hy2.example.com", Port: 443, Password: "pass", Params: map[string]string{"sni": "example.com"}},
		protocol.AnyTLSNode{Base: protocol.Base{Type: "anytls", Name: "AnyTLS"}, Host: "anytls.example.com", Port: 443, Password: "pass", Params: map[string]string{"sni": "example.com"}},
		protocol.TUICNode{Base: protocol.Base{Type: "tuic", Name: "TUIC"}, Host: "tuic.example.com", Port: 443, UUID: "uuid", Password: "pass", Params: map[string]string{"sni": "example.com"}},
	}

	var out strings.Builder
	if err := builtintemplate.RenderBuiltin(&out, "mihomo", nodes); err != nil {
		t.Fatalf("template RenderBuiltin returned error: %v", err)
	}
	var decoded map[string]any
	if err := yaml.Unmarshal([]byte(out.String()), &decoded); err != nil {
		t.Fatalf("rendered invalid YAML: %v\n%s", err, out.String())
	}
	proxies := decoded["proxies"].([]any)
	if len(proxies) != len(nodes) {
		t.Fatalf("got %d proxies, want %d", len(proxies), len(nodes))
	}
	for i, want := range []string{"ssr", "vmess", "trojan", "hysteria2", "anytls", "tuic"} {
		if got := proxies[i].(map[string]any)["type"]; got != want {
			t.Errorf("proxy %d type = %q, want %q", i, got, want)
		}
	}
}
