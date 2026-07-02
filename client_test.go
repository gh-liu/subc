package subc

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestFetchAndParseUsesClientRequestAndParser(t *testing.T) {
	client := ShadowrocketClient{}
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if got := r.Header.Get("User-Agent"); got != DefaultUserAgent {
			t.Fatalf("User-Agent = %q, want %q", got, DefaultUserAgent)
		}
		if got := r.Header.Get("Accept"); got != "*/*" {
			t.Fatalf("Accept = %q, want */*", got)
		}
		_, _ = w.Write([]byte("STATUS=ok\ntrojan://password@example.com:443#Node"))
	}))
	defer server.Close()

	nodes, err := FetchAndParse(server.URL, client)
	if err != nil {
		t.Fatalf("FetchAndParse returned error: %v", err)
	}
	if len(nodes) != 1 || nodes[0].Protocol() != "trojan" || nodes[0].NodeName() != "Node" {
		t.Fatalf("unexpected nodes: %+v", nodes)
	}
}
