package protocol

import "testing"

func TestParseTUIC(t *testing.T) {
	node, err := Parse("tuic://uuid:pass@example.com:443?congestion_control=bbr&sni=example.org&allowInsecure=1#TUIC%20Node")
	if err != nil {
		t.Fatalf("Parse returned error: %v", err)
	}
	n, ok := node.(TUICNode)
	if !ok {
		t.Fatalf("got %T, want TUICNode", node)
	}
	if n.Protocol() != "tuic" || n.NodeName() != "TUIC Node" || n.Host != "example.com" || n.Port != 443 || n.UUID != "uuid" || n.Password != "pass" {
		t.Fatalf("unexpected tuic node: %+v", n)
	}
	if n.Params["congestion_control"] != "bbr" || n.Params["sni"] != "example.org" || n.Params["allowInsecure"] != "1" {
		t.Fatalf("unexpected tuic params: %+v", n.Params)
	}
}
