package protocol

import "strings"

type VLESSNode struct {
	Base
	Host   string            `json:"host,omitempty"`
	Port   int               `json:"port,omitempty"`
	UUID   string            `json:"uuid,omitempty"`
	Params map[string]string `json:"params,omitempty"`
}

func init() { Register("vless", ParseVLESS) }

func ParseVLESS(raw string) (Node, error) {
	base, userinfo, host, port, params, err := parseUserInfo(raw, "vless")
	if err != nil {
		return nil, err
	}
	return VLESSNode{Base: base, Host: host, Port: port, UUID: strings.TrimPrefix(userinfo, "auto:"), Params: params}, nil
}
