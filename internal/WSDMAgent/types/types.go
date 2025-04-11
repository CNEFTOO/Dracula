package types

import (
	"net/http"
	"net/url"
)

type DNSTypes struct {
	Name     string `json:"name"`
	Type     string `json:"type"`
	SourceIP string `json:"source_ip"`
	DestIP   string `json:"dest_ip"`
}

func NewDnsTypes(name, t, source, dest string) *DNSTypes {
	return &DNSTypes{Name: name, Type: t, SourceIP: source, DestIP: dest}
}

type HttpRequest struct {
	Host              string
	IP                string
	Client            string
	Port              string
	URL               *url.URL
	Header            http.Header
	RequestUri        string
	RequestBody       string
	Method            string
	RequestParameters url.Values
}

func NewHttpRequest(request *http.Request, client, ip, port string) (*HttpRequest, error) {
	return nil, nil
}

type TCPConnectionInfo struct {
	Protocol   string `json:"protocol"`
	SourceIP   string `json:"source_ip"`
	DestIP     string `json:"dest_ip"`
	SourcePort string `json:"source_port"`
	DestPort   string `json:"dest_port"`
}

func NewTCPConnectionInfo(protocol, sourceIP, destIP, sourcePort, destPort string) *TCPConnectionInfo {
	return &TCPConnectionInfo{
		Protocol:   protocol,
		SourceIP:   sourceIP,
		DestIP:     destIP,
		SourcePort: sourcePort,
		DestPort:   destPort,
	}
}
