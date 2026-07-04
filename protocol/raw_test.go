package protocol

import "testing"

func TestParseUnknownProtocolReturnsRawNode(t *testing.T) {
	node, err := Parse("unknown://example")
	if err != nil {
		t.Fatalf("Parse returned error: %v", err)
	}
	if node.Protocol() != "unknown" || node.RawURI() != "unknown://example" {
		t.Fatalf("unexpected raw node: %+v", node)
	}
}
