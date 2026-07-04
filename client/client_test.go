package client

import "testing"

func TestGetDefaultClient(t *testing.T) {
	c, err := Get("")
	if err != nil {
		t.Fatalf("Get returned error: %v", err)
	}
	if _, ok := c.(ShadowrocketClient); !ok {
		t.Fatalf("got %T, want ShadowrocketClient", c)
	}
}

func TestGetV2RayClient(t *testing.T) {
	c, err := Get("v2ray")
	if err != nil {
		t.Fatalf("Get returned error: %v", err)
	}
	if _, ok := c.(V2RayClient); !ok {
		t.Fatalf("got %T, want V2RayClient", c)
	}
}

func TestGetUnknownClient(t *testing.T) {
	if _, err := Get("unknown"); err == nil {
		t.Fatal("Get returned nil error")
	}
}
