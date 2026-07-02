package protocol

import (
	"errors"
	"net/url"
	"strings"
)

type ShadowsocksNode struct {
	Base
	Host     string `json:"host,omitempty"`
	Port     int    `json:"port,omitempty"`
	Method   string `json:"method,omitempty"`
	Password string `json:"password,omitempty"`
}

func init() { Register("ss", ParseShadowsocks) }

func ParseShadowsocks(raw string) (Node, error) {
	u, err := url.Parse(raw)
	if err != nil {
		return nil, err
	}
	n := ShadowsocksNode{Base: Base{Type: "ss", Name: fragmentName(u), Raw: raw}}
	encoded := strings.TrimPrefix(raw, "ss://")
	if hash := strings.LastIndex(encoded, "#"); hash >= 0 {
		encoded = encoded[:hash]
	}
	if q := strings.LastIndex(encoded, "?"); q >= 0 {
		encoded = encoded[:q]
	}

	var userHost string
	if strings.Contains(encoded, "@") {
		parts := strings.SplitN(encoded, "@", 2)
		userinfo, ok := decodeBase64Text(parts[0])
		if !ok {
			userinfo = parts[0]
		}
		userHost = userinfo + "@" + parts[1]
	} else {
		decoded, ok := decodeBase64Text(encoded)
		if !ok {
			return nil, errors.New("invalid ss base64 payload")
		}
		userHost = decoded
	}

	userinfo, hostport, ok := strings.Cut(userHost, "@")
	if !ok {
		return nil, errors.New("invalid ss payload")
	}
	n.Method, n.Password, _ = strings.Cut(userinfo, ":")
	n.Host, n.Port, err = setHostPort(hostport)
	if err != nil {
		return nil, err
	}
	return n, nil
}
