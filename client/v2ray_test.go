package client

import (
	"net/http"
	"strings"
	"testing"
)

func TestV2RayClientPreparesRequest(t *testing.T) {
	req, err := http.NewRequest(http.MethodGet, "https://example.com/sub", nil)
	if err != nil {
		t.Fatal(err)
	}

	V2RayClient{}.PrepareRequest(req)

	if got := req.Header.Get("User-Agent"); got != DefaultV2RayUserAgent {
		t.Fatalf("User-Agent = %q, want %q", got, DefaultV2RayUserAgent)
	}
	if got := req.Header.Get("Accept"); got != "*/*" {
		t.Fatalf("Accept = %q, want */*", got)
	}
}

func TestV2RayClientParsesLineOrientedSubscription(t *testing.T) {
	nodes, err := V2RayClient{}.Parse(strings.NewReader("STATUS=ok\ntrojan://password@example.com:443#Node"))
	if err != nil {
		t.Fatalf("Parse returned error: %v", err)
	}
	if len(nodes) != 1 || nodes[0].Protocol() != "trojan" || nodes[0].NodeName() != "Node" {
		t.Fatalf("unexpected nodes: %+v", nodes)
	}
}
