package protocol

type TrojanNode struct {
	Base
	Host     string            `json:"host,omitempty"`
	Port     int               `json:"port,omitempty"`
	Password string            `json:"password,omitempty"`
	Params   map[string]string `json:"params,omitempty"`
}

func init() { Register("trojan", ParseTrojan) }

func ParseTrojan(raw string) (Node, error) {
	base, userinfo, host, port, params, err := parseUserInfo(raw, "trojan")
	if err != nil {
		return nil, err
	}
	return TrojanNode{Base: base, Host: host, Port: port, Password: userinfo, Params: params}, nil
}
