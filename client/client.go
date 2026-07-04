package client

import (
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/gh-liu/subc/protocol"
)

// Parser parses a subscription response body into proxy nodes.
type Parser interface {
	Parse(r io.Reader) ([]protocol.Node, error)
}

// Client describes a subscription client profile. It prepares requests and
// parses the response format returned for that client.
type Client interface {
	Parser
	PrepareRequest(req *http.Request)
}

type Factory func() Client

var registry = map[string]Factory{}

func Register(name string, factory Factory) {
	registry[strings.ToLower(name)] = factory
}

func Get(name string) (Client, error) {
	if name == "" {
		name = "shadowrocket"
	}
	factory, ok := registry[strings.ToLower(name)]
	if !ok {
		return nil, fmt.Errorf("unsupported client profile %q", name)
	}
	return factory(), nil
}
