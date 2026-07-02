package protocol

import "testing"

func TestParseVMess(t *testing.T) {
	node, err := Parse("vmess://eyJ2IjoiMiIsInBzIjoiVk1lc3MgTm9kZSIsImFkZCI6ImV4YW1wbGUuY29tIiwicG9ydCI6IjEwMDAiLCJpZCI6InV1aWQiLCJhaWQiOiIwIiwibmV0Ijoid3MiLCJ0eXBlIjoibm9uZSIsImhvc3QiOiJ3cy5leGFtcGxlLmNvbSIsInBhdGgiOiIvY2hhdCIsInRscyI6InRscyJ9")
	if err != nil {
		t.Fatalf("Parse returned error: %v", err)
	}
	n, ok := node.(VMessNode)
	if !ok {
		t.Fatalf("got %T, want VMessNode", node)
	}
	if n.Protocol() != "vmess" || n.NodeName() != "VMess Node" || n.Host != "example.com" || n.Port != 1000 || n.UUID != "uuid" || n.Params["net"] != "ws" || n.Params["path"] != "/chat" {
		t.Fatalf("unexpected vmess node: %+v", n)
	}
}
