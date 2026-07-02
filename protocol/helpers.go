package protocol

import (
	"encoding/base64"
	"net/url"
	"strconv"
	"strings"
)

func DecodeBase64Text(s string) (string, bool) {
	s = strings.TrimSpace(s)
	for _, enc := range []*base64.Encoding{base64.StdEncoding, base64.RawStdEncoding, base64.URLEncoding, base64.RawURLEncoding} {
		if b, err := enc.DecodeString(s); err == nil {
			return string(b), true
		}
	}
	if m := len(s) % 4; m != 0 {
		return DecodeBase64Text(s + strings.Repeat("=", 4-m))
	}
	return "", false
}

func decodeBase64Text(s string) (string, bool) {
	return DecodeBase64Text(s)
}

func setHostPort(hostport string) (string, int, error) {
	u, err := url.Parse("scheme://" + hostport)
	if err != nil {
		return "", 0, err
	}
	port, _ := strconv.Atoi(u.Port())
	return u.Hostname(), port, nil
}

func fragmentName(u *url.URL) string {
	name, _ := url.QueryUnescape(u.Fragment)
	return name
}

func stringValue(v any) string {
	switch x := v.(type) {
	case string:
		return x
	case float64:
		return strconv.FormatFloat(x, 'f', -1, 64)
	default:
		return ""
	}
}
