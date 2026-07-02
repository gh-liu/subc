package protocol

type AnyTLSNode struct {
	Base
	Host     string            `json:"host,omitempty"`
	Port     int               `json:"port,omitempty"`
	Password string            `json:"password,omitempty"`
	Params   map[string]string `json:"params,omitempty"`
}

func init() { Register("anytls", ParseAnyTLS) }

func ParseAnyTLS(raw string) (Node, error) {
	base, userinfo, host, port, params, err := parseUserInfo(raw, "anytls")
	if err != nil {
		return nil, err
	}
	return AnyTLSNode{Base: base, Host: host, Port: port, Password: userinfo, Params: params}, nil
}
