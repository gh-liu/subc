package protocol

import (
	"encoding/json"
	"errors"
	"strconv"
	"strings"
)

type VMessNode struct {
	Base
	Host   string            `json:"host,omitempty"`
	Port   int               `json:"port,omitempty"`
	UUID   string            `json:"uuid,omitempty"`
	Params map[string]string `json:"params,omitempty"`
}

func init() { Register("vmess", ParseVMess) }

func ParseVMess(raw string) (Node, error) {
	payload := strings.TrimPrefix(raw, "vmess://")
	decoded, ok := decodeBase64Text(payload)
	if !ok {
		return nil, errors.New("invalid vmess base64 payload")
	}
	var data map[string]any
	if err := json.Unmarshal([]byte(decoded), &data); err != nil {
		return nil, err
	}
	n := VMessNode{Base: Base{Type: "vmess", Name: stringValue(data["ps"]), Raw: raw}, Host: stringValue(data["add"]), UUID: stringValue(data["id"]), Params: map[string]string{}}
	if port, err := strconv.Atoi(stringValue(data["port"])); err == nil {
		n.Port = port
	}
	for _, key := range []string{"v", "aid", "net", "type", "host", "path", "tls", "sni", "alpn", "fp", "scy"} {
		if value := stringValue(data[key]); value != "" {
			n.Params[key] = value
		}
	}
	return n, nil
}
