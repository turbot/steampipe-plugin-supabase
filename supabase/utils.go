package supabase

import (
	"context"
	"encoding/json"
	"net"
	"net/http"
	"time"
)

const (
	// Ref: https://www.iana.org/assignments/dns-parameters/dns-parameters.xhtml#dns-parameters-4
	dnsIPv4Type uint16 = 1
	cnameType   uint16 = 5
	dnsIPv6Type uint16 = 28
)

type dnsAnswer struct {
	Type uint16 `json:"type"`
	Data string `json:"data"`
}

type dnsResponse struct {
	Answer []dnsAnswer `json:",omitempty"`
}

// Performs DNS lookup via HTTPS, in case firewall blocks native netgo resolver.
func fallbackLookupIP(ctx context.Context, address string) string {
	host, port, err := net.SplitHostPort(address)
	if err != nil {
		return ""
	}
	// Ref: https://developers.cloudflare.com/1.1.1.1/encryption/dns-over-https/make-api-requests/dns-json
	req, err := http.NewRequestWithContext(ctx, "GET", "https://1.1.1.1/dns-query?name="+host, nil)
	if err != nil {
		return ""
	}
	req.Header.Add("accept", "application/dns-json")

	httpClient := http.Client{Timeout: 10 * time.Second}
	// Sends request
	resp, err := httpClient.Do(req)
	if err != nil {
		return ""
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return ""
	}
	// Parses response
	var data dnsResponse
	dec := json.NewDecoder(resp.Body)
	if err := dec.Decode(&data); err != nil {
		return ""
	}
	// Look for first valid IP
	for _, answer := range data.Answer {
		if answer.Type == dnsIPv4Type || answer.Type == dnsIPv6Type {
			return net.JoinHostPort(answer.Data, port)
		}
	}
	return ""
}
