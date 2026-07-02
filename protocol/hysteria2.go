package protocol

type Hysteria2Node struct {
	Base
	Host     string            `json:"host,omitempty"`
	Port     int               `json:"port,omitempty"`
	Password string            `json:"password,omitempty"`
	Params   map[string]string `json:"params,omitempty"`
}

func init() {
	Register("hysteria2", ParseHysteria2)
	Register("hy2", ParseHysteria2)
}

func ParseHysteria2(raw string) (Node, error) {
	base, userinfo, host, port, params, err := parseUserInfo(raw, "hysteria2")
	if err != nil {
		return nil, err
	}
	return Hysteria2Node{Base: base, Host: host, Port: port, Password: userinfo, Params: params}, nil
}
