package client

import (
	"net/http"

	"github.com/gh-liu/subc/protocol"
)

const DefaultUserAgent = "Shadowrocket"

func init() {
	Register("shadowrocket", func() Client { return ShadowrocketClient{} })
}

// ShadowrocketClient requests and parses Shadowrocket-compatible subscriptions.
type ShadowrocketClient struct {
	URIListClient
}

func (c ShadowrocketClient) PrepareRequest(req *http.Request) {
	c.URIListClient.UserAgent = DefaultUserAgent
	c.URIListClient.PrepareRequest(req)
}

// ParseShadowrocket parses a Shadowrocket subscription body. It accepts
// base64-wrapped line-oriented URI lists and plain line-oriented URI lists.
func ParseShadowrocket(content string) ([]protocol.Node, error) {
	return ParseURIListSubscription(content)
}
