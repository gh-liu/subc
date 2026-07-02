package protocol

import (
	"errors"
	"net/url"
	"strconv"
	"strings"
)

type SSRNode struct {
	Base
	Host     string            `json:"host,omitempty"`
	Port     int               `json:"port,omitempty"`
	Method   string            `json:"method,omitempty"`
	Password string            `json:"password,omitempty"`
	Params   map[string]string `json:"params,omitempty"`
}

func init() { Register("ssr", ParseSSR) }

func ParseSSR(raw string) (Node, error) {
	payload := strings.TrimPrefix(raw, "ssr://")
	decoded, ok := decodeBase64Text(payload)
	if !ok {
		return nil, errors.New("invalid ssr base64 payload")
	}
	mainPart, queryPart, _ := strings.Cut(decoded, "/?")
	parts := strings.Split(mainPart, ":")
	if len(parts) != 6 {
		return nil, errors.New("invalid ssr payload")
	}
	n := SSRNode{Base: Base{Type: "ssr", Raw: raw}, Host: parts[0], Method: parts[3], Params: map[string]string{"protocol": parts[2], "obfs": parts[4]}}
	if port, err := strconv.Atoi(parts[1]); err == nil {
		n.Port = port
	}
	if password, ok := decodeBase64Text(parts[5]); ok {
		n.Password = password
	}
	values, err := url.ParseQuery(queryPart)
	if err != nil {
		return nil, err
	}
	for key, vals := range values {
		if len(vals) == 0 {
			continue
		}
		value := vals[0]
		if decoded, ok := decodeBase64Text(value); ok {
			value = decoded
		}
		if key == "remarks" {
			n.Name = value
			continue
		}
		n.Params[key] = value
	}
	return n, nil
}
