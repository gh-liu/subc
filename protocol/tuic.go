package protocol

import "strings"

type TUICNode struct {
	Base
	Host     string            `json:"host,omitempty"`
	Port     int               `json:"port,omitempty"`
	UUID     string            `json:"uuid,omitempty"`
	Password string            `json:"password,omitempty"`
	Params   map[string]string `json:"params,omitempty"`
}

func init() { Register("tuic", ParseTUIC) }

func ParseTUIC(raw string) (Node, error) {
	base, userinfo, host, port, params, err := parseUserInfo(raw, "tuic")
	if err != nil {
		return nil, err
	}
	uuid, password, _ := strings.Cut(userinfo, ":")
	return TUICNode{Base: base, Host: host, Port: port, UUID: uuid, Password: password, Params: params}, nil
}
