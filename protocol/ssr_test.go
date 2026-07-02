package protocol

import "testing"

func TestParseSSR(t *testing.T) {
	raw := "ssr://ZXhhbXBsZS5jb206MTIzNDpvcmlnaW46YWVzLTI1Ni1jZmI6cGxhaW46Y0dGemN3Lz9yZW1hcmtzPVUxTlNJRTV2WkdVJm9iZnNwYXJhbT0mcHJvdG9wYXJhbT0"
	node, err := Parse(raw)
	if err != nil {
		t.Fatalf("Parse returned error: %v", err)
	}
	n, ok := node.(SSRNode)
	if !ok {
		t.Fatalf("got %T, want SSRNode", node)
	}
	if n.Protocol() != "ssr" || n.NodeName() != "SSR Node" || n.Host != "example.com" || n.Port != 1234 || n.Method != "aes-256-cfb" || n.Password != "pass" {
		t.Fatalf("unexpected ssr node: %+v", n)
	}
	if n.Params["protocol"] != "origin" || n.Params["obfs"] != "plain" {
		t.Fatalf("unexpected ssr params: %+v", n.Params)
	}
}
