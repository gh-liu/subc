package subc

import (
	"io"
	"net/http"
)

// Parser parses a subscription response body into proxy nodes.
type Parser interface {
	Parse(r io.Reader) ([]Node, error)
}

// Client describes a subscription client profile. It prepares requests and
// parses the response format returned for that client.
type Client interface {
	Parser
	PrepareRequest(req *http.Request)
}
