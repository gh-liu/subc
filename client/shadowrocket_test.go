package client

import "testing"

func TestParseSkipsSubscriptionStatusLines(t *testing.T) {
	content := "STATUS=🚀:49.39GB,↓:94.59GB,TOT:200GB⚡Expires:2026-10-09\nss://YWVzLTI1Ni1nY206cGFzc0BleGFtcGxlLmNvbToxMjM0#SS%20Node"

	nodes, err := ParseShadowrocket(content)
	if err != nil {
		t.Fatalf("Parse returned error: %v", err)
	}
	if len(nodes) != 1 {
		t.Fatalf("got %d nodes, want 1", len(nodes))
	}
	if nodes[0].Protocol() != "ss" || nodes[0].NodeName() != "SS Node" {
		t.Fatalf("unexpected node: %+v", nodes[0])
	}
}

func TestParseShadowrocketSplitsByLineNotWhitespace(t *testing.T) {
	content := "unknown://example.com#Name With Spaces\nss://YWVzLTI1Ni1nY206cGFzc0BleGFtcGxlLmNvbToxMjM0#SS%20Node"

	nodes, err := ParseShadowrocket(content)
	if err != nil {
		t.Fatalf("Parse returned error: %v", err)
	}
	if len(nodes) != 2 {
		t.Fatalf("got %d nodes, want 2", len(nodes))
	}
	if nodes[0].RawURI() != "unknown://example.com#Name With Spaces" {
		t.Fatalf("unexpected raw URI: %q", nodes[0].RawURI())
	}
}
