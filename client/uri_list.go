package client

import (
	"io"
	"net/http"
	"strings"

	"github.com/gh-liu/subc/protocol"
)

type URIListClient struct {
	UserAgent string
}

func (c URIListClient) PrepareRequest(req *http.Request) {
	if c.UserAgent != "" {
		req.Header.Set("User-Agent", c.UserAgent)
	}
	req.Header.Set("Accept", "*/*")
}

func (URIListClient) Parse(r io.Reader) ([]protocol.Node, error) {
	body, err := io.ReadAll(r)
	if err != nil {
		return nil, err
	}
	return ParseURIListSubscription(string(body))
}

// ParseURIListSubscription parses a base64-wrapped or plain line-oriented URI list.
func ParseURIListSubscription(content string) ([]protocol.Node, error) {
	content = strings.TrimSpace(content)
	if content == "" {
		return nil, nil
	}
	if decoded, ok := protocol.DecodeBase64Text(content); ok && strings.Contains(decoded, "://") {
		content = decoded
	}

	var nodes []protocol.Node
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
