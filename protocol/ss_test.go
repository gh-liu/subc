package protocol

import "testing"

func TestParseShadowsocks(t *testing.T) {
	node, err := Parse("ss://YWVzLTI1Ni1nY206cGFzc0BleGFtcGxlLmNvbToxMjM0#SS%20Node")
	if err != nil {
		t.Fatalf("Parse returned error: %v", err)
	}
	n, ok := node.(ShadowsocksNode)
	if !ok {
		t.Fatalf("got %T, want ShadowsocksNode", node)
	}
	if n.Protocol() != "ss" || n.NodeName() != "SS Node" || n.Host != "example.com" || n.Port != 1234 || n.Method != "aes-256-gcm" || n.Password != "pass" {
		t.Fatalf("unexpected ss node: %+v", n)
	}
}
