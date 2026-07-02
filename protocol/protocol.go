package protocol

import (
	"fmt"
	"strings"
)

// Node is implemented by every parsed proxy node type.
type Node interface {
	Protocol() string
	NodeName() string
	RawURI() string
}

type Parser func(raw string) (Node, error)

var parsers = map[string]Parser{}

func Register(name string, parser Parser) {
	parsers[strings.ToLower(name)] = parser
}

func Parse(raw string) (Node, error) {
	schemeEnd := strings.Index(raw, "://")
	if schemeEnd <= 0 {
		return nil, fmt.Errorf("invalid node URI: %q", raw)
	}
	name := strings.ToLower(raw[:schemeEnd])
	parser, ok := parsers[name]
	if !ok {
		return RawNode{Base: Base{Type: name, Raw: raw}}, nil
	}
	return parser(raw)
}

type Base struct {
	Type string `json:"protocol"`
	Name string `json:"name,omitempty"`
	Raw  string `json:"raw"`
}

func (n Base) Protocol() string { return n.Type }
func (n Base) NodeName() string { return n.Name }
func (n Base) RawURI() string   { return n.Raw }

type RawNode struct {
	Base
}
