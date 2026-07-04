package protocol

import "testing"

func TestParseVLESS(t *testing.T) {
	node, err := Parse("vless://uuid@example.com:443?encryption=none&security=tls&type=ws&path=%2Fws#VLESS%20Node")
	if err != nil {
		t.Fatalf("Parse returned error: %v", err)
	}
	n, ok := node.(VLESSNode)
	if !ok {
		t.Fatalf("got %T, want VLESSNode", node)
	}
	if n.Protocol() != "vless" || n.NodeName() != "VLESS Node" || n.Host != "example.com" || n.Port != 443 || n.UUID != "uuid" || n.Params["type"] != "ws" || n.Params["path"] != "/ws" {
		t.Fatalf("unexpected vless node: %+v", n)
	}
}

func TestParseVLESSWithBase64Authority(t *testing.T) {
	raw := "vless://YXV0bzo2ZjBmNjdkYi1kZjdlLTQ3ZDktYjQ0Ny1hMjMyZTQ0MzNiNTNAaGszLm1peWF6b25vLWthb3JpLmNvbTo0NDM=?tfo=1&remark=%F0%9F%87%AD%F0%9F%87%B0Hong%20Kong%2003&tls=1&xtls=2&sni=hk4e-launcher-static.hoyoverse.com&pbk=zV4Sja0iajpdzkR1iBMbI7RRWL2nlpgwGOtirQSG-zc&sid=b81edf7fdee4&fp=chrome"
	node, err := Parse(raw)
	if err != nil {
		t.Fatalf("Parse returned error: %v", err)
	}
	n, ok := node.(VLESSNode)
	if !ok {
		t.Fatalf("got %T, want VLESSNode", node)
	}
	if n.UUID != "6f0f67db-df7e-47d9-b447-a232e4433b53" || n.Host != "hk3.miyazono-kaori.com" || n.Port != 443 {
		t.Fatalf("unexpected vless endpoint: %+v", n)
	}
	if n.Params["sni"] != "hk4e-launcher-static.hoyoverse.com" || n.Params["remark"] != "🇭🇰Hong Kong 03" {
		t.Fatalf("unexpected vless params: %+v", n.Params)
	}
	if n.NodeName() != "🇭🇰Hong Kong 03" {
		t.Fatalf("unexpected vless node name: %q", n.NodeName())
	}
}
