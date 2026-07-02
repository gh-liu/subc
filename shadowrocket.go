package subc

import (
	"io"
	"net/http"
	"strings"

	"github.com/gh-liu/subc/protocol"
)

const DefaultUserAgent = "Shadowrocket"

// ShadowrocketClient requests and parses Shadowrocket-compatible subscriptions.
type ShadowrocketClient struct{}

func (ShadowrocketClient) PrepareRequest(req *http.Request) {
	req.Header.Set("User-Agent", DefaultUserAgent)
	req.Header.Set("Accept", "*/*")
}

func (ShadowrocketClient) Parse(r io.Reader) ([]Node, error) {
	body, err := io.ReadAll(r)
	if err != nil {
		return nil, err
	}
	return parseShadowrocket(string(body))
}

// parseShadowrocket parses a Shadowrocket subscription body. It accepts
// base64-wrapped line-oriented URI lists and plain line-oriented URI lists.
func parseShadowrocket(content string) ([]Node, error) {
	content = strings.TrimSpace(content)
	if content == "" {
		return nil, nil
	}
	if decoded, ok := protocol.DecodeBase64Text(content); ok && strings.Contains(decoded, "://") {
		content = decoded
	}

	var nodes []Node
	for _, line := range strings.Split(content, "\n") {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		if !strings.Contains(line, "://") {
			continue
		}
		node, err := protocol.Parse(line)
		if err != nil {
			return nil, err
		}
		nodes = append(nodes, node)
	}
	return nodes, nil
}
