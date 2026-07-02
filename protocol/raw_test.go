package protocol

import "testing"

func TestParseUnknownProtocolReturnsRawNode(t *testing.T) {
	node, err := Parse("tuic://example")
	if err != nil {
		t.Fatalf("Parse returned error: %v", err)
	}
	if node.Protocol() != "tuic" || node.RawURI() != "tuic://example" {
		t.Fatalf("unexpected raw node: %+v", node)
	}
}
