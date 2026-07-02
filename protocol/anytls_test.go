package protocol

import "testing"

func TestParseAnyTLS(t *testing.T) {
	raw := "anytls://bb502703-bce0-4e41-a885-32d4f5decfc0@7jp.cloudfrontcdn.com:4430?allowInsecure=1&peer=dss0.bdstatic.com&udp=1&tfo=0#%F0%9F%87%AF%F0%9F%87%B57%E6%97%A5%E6%9C%AC-%E8%81%94%E9%80%9A%2F%E7%A7%BB%E5%8A%A8%28AnyTLS%29"
	node, err := Parse(raw)
	if err != nil {
		t.Fatalf("Parse returned error: %v", err)
	}
	n, ok := node.(AnyTLSNode)
	if !ok {
		t.Fatalf("got %T, want AnyTLSNode", node)
	}
	if n.Protocol() != "anytls" || n.Password != "bb502703-bce0-4e41-a885-32d4f5decfc0" || n.Host != "7jp.cloudfrontcdn.com" || n.Port != 4430 || n.NodeName() != "🇯🇵7日本-联通/移动(AnyTLS)" {
		t.Fatalf("unexpected anytls node: %+v", n)
	}
	if n.Params["peer"] != "dss0.bdstatic.com" || n.Params["allowInsecure"] != "1" || n.Params["tfo"] != "0" {
		t.Fatalf("unexpected anytls params: %+v", n.Params)
	}
}
