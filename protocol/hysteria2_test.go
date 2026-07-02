package protocol

import "testing"

func TestParseHysteria2(t *testing.T) {
	raw := "hysteria2://bb502703-bce0-4e41-a885-32d4f5decfc0@36ph.cloudfrontcdn.com:4433?insecure=1&sni=&tfo=1&udp=0&mport=40000-50000#%F0%9F%87%B5%F0%9F%87%AD36%E8%8F%B2%E5%BE%8B%E5%AE%BE-%E5%85%A8%E7%BD%91%E4%BC%98%E5%8C%96%28hy2%29"
	node, err := Parse(raw)
	if err != nil {
		t.Fatalf("Parse returned error: %v", err)
	}
	n, ok := node.(Hysteria2Node)
	if !ok {
		t.Fatalf("got %T, want Hysteria2Node", node)
	}
	if n.Protocol() != "hysteria2" || n.Password != "bb502703-bce0-4e41-a885-32d4f5decfc0" || n.Host != "36ph.cloudfrontcdn.com" || n.Port != 4433 || n.NodeName() != "🇵🇭36菲律宾-全网优化(hy2)" {
		t.Fatalf("unexpected hysteria2 node: %+v", n)
	}
	if n.Params["mport"] != "40000-50000" || n.Params["tfo"] != "1" || n.Params["udp"] != "0" {
		t.Fatalf("unexpected hysteria2 params: %+v", n.Params)
	}
}
