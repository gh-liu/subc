package protocol

import (
	"net/url"
	"strings"
)

func parseUserInfo(raw, protocolName string) (Base, string, string, int, map[string]string, error) {
	u, err := url.Parse(raw)
	if err != nil {
		return Base{}, "", "", 0, nil, err
	}
	userinfo := u.User.Username()
	if password, ok := u.User.Password(); ok {
		userinfo += ":" + password
	}
	hostport := u.Host
	if decoded, ok := decodeBase64Text(u.Host); ok && strings.Contains(decoded, "@") {
		userinfo, hostport, _ = strings.Cut(decoded, "@")
	}
	host, port, err := setHostPort(hostport)
	if err != nil {
		return Base{}, "", "", 0, nil, err
	}
	params := map[string]string{}
	for key, values := range u.Query() {
		if len(values) > 0 {
			params[key] = values[0]
		}
	}
	name := fragmentName(u)
	if name == "" {
		name = params["remark"]
	}
	return Base{Type: protocolName, Name: name, Raw: raw}, userinfo, host, port, params, nil
}
