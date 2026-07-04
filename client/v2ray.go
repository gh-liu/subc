package client

import "net/http"

const DefaultV2RayUserAgent = "v2rayN"

func init() {
	Register("v2ray", func() Client { return V2RayClient{} })
}

// V2RayClient requests and parses V2Ray-compatible subscriptions.
type V2RayClient struct {
	URIListClient
}

func (c V2RayClient) PrepareRequest(req *http.Request) {
	c.URIListClient.UserAgent = DefaultV2RayUserAgent
	c.URIListClient.PrepareRequest(req)
}
