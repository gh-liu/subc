package subc

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gh-liu/subc/protocol"
)

// Node is implemented by every parsed proxy node type.
type Node = protocol.Node

var defaultHTTPClient = &http.Client{Timeout: 30 * time.Second}

// FetchAndParse downloads a subscription URL and parses all nodes with client.
func FetchAndParse(subscriptionURL string, client Client) ([]Node, error) {
	req, err := http.NewRequest(http.MethodGet, subscriptionURL, nil)
	if err != nil {
		return nil, err
	}
	client.PrepareRequest(req)

	resp, err := defaultHTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return nil, fmt.Errorf("fetch subscription: status %s", resp.Status)
	}
	return client.Parse(resp.Body)
}
