package wsdmagent

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/url"
	"strings"

	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"github.com/google/gopacket/tcpassembly"
	"github.com/google/gopacket/tcpassembly/tcpreader"
	"github.com/seaung/Dracula/internal/WSDMAgent/types"
)

type httpStreamFactory struct{}

type httpStream struct {
	net, transport gopacket.Flow
	r              tcpreader.ReaderStream
}

func (h *httpStreamFactory) New(net, transport gopacket.Flow) tcpassembly.Stream {
	hStream := &httpStream{
		net:       net,
		transport: transport,
		r:         tcpreader.NewReaderStream(),
	}

	go hStream.run()

	return &hStream.r
}

func (h *httpStream) run() {}

// 处理数据包
func processPacket(packet gopacket.Packet) {
	ipLayer := packet.Layer(layers.LayerTypeIPv4)
	if ipLayer != nil {
		ip, _ := ipLayer.(*layers.IPv4)
		if ip != nil {
			switch ip.Protocol {
			case layers.IPProtocolTCP:
				// 处理 TCP 数据包
				tcpLayer := packet.Layer(layers.LayerTypeTCP)
				if tcpLayer != nil {
					tcp, _ := tcpLayer.(*layers.TCP)
					if tcp != nil {
						// 创建TCP连接信息
						tcpInfo := types.NewTCPConnectionInfo(
							"TCP",
							ip.SrcIP.String(),
							ip.DstIP.String(),
							tcp.SrcPort.String(),
							tcp.DstPort.String(),
						)
						// 发送TCP连接信息
						err := sendPacket(tcpInfo)
						if err != nil {
							// 处理错误，可以记录日志
						}
					}
				}
			case layers.IPProtocolUDP:
				// 处理 UDP 数据包
				udpLayer := packet.Layer(layers.LayerTypeUDP)
				if udpLayer != nil {
					udp, _ := udpLayer.(*layers.UDP)
					if udp != nil {
						// 创建UDP连接信息，复用TCP连接信息结构体
						udpInfo := types.NewTCPConnectionInfo(
							"UDP",
							ip.SrcIP.String(),
							ip.DstIP.String(),
							udp.SrcPort.String(),
							udp.DstPort.String(),
						)
						// 发送UDP连接信息
						err := sendPacket(udpInfo)
						if err != nil {
							// 处理错误，可以记录日志
						}

						// 检查是否为DNS请求
						if udp.SrcPort == 53 || udp.DstPort == 53 {
							parseDns(packet)
						}
					}
				}
			case layers.IPProtocolICMPv4:
				// 处理 ICMP 数据包
				icmpLayer := packet.Layer(layers.LayerTypeICMPv4)
				if icmpLayer != nil {
					icmp, _ := icmpLayer.(*layers.ICMPv4)
					if icmp != nil {
						// 创建ICMP连接信息，复用TCP连接信息结构体
						icmpInfo := types.NewTCPConnectionInfo(
							"ICMPv4",
							ip.SrcIP.String(),
							ip.DstIP.String(),
							"", // ICMP没有端口
							"", // ICMP没有端口
						)
						// 发送ICMP连接信息
						err := sendPacket(icmpInfo)
						if err != nil {
							// 处理错误，可以记录日志
						}
					}
				}
			case layers.IPProtocolICMPv6:
				// 处理 ICMPv6 数据包
				icmpv6Layer := packet.Layer(layers.LayerTypeICMPv6)
				if icmpv6Layer != nil {
					icmpv6, _ := icmpv6Layer.(*layers.ICMPv6)
					if icmpv6 != nil {
						// 创建ICMPv6连接信息，复用TCP连接信息结构体
						icmpv6Info := types.NewTCPConnectionInfo(
							"ICMPv6",
							ip.SrcIP.String(),
							ip.DstIP.String(),
							"", // ICMPv6没有端口
							"", // ICMPv6没有端口
						)
						// 发送ICMPv6连接信息
						err := sendPacket(icmpv6Info)
						if err != nil {
							// 处理错误，可以记录日志
						}
					}
				}
			}
		}
	}
}

func sendPacket(tcpConnectionInfo *types.TCPConnectionInfo) error {
	data, err := json.Marshal(tcpConnectionInfo)
	if err != nil {
		return err
	}

	resp, err := http.Post("http://127.0.0.1:8080/api/v1/wsdmagent/tcp", "application/json", bytes.NewBuffer(data))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return nil
}

func parseDns(packet gopacket.Packet) {
	// 获取DNS层
	dnsLayer := packet.Layer(layers.LayerTypeDNS)
	if dnsLayer == nil {
		return
	}

	// 解析DNS数据
	dns, ok := dnsLayer.(*layers.DNS)
	if !ok || dns == nil {
		return
	}

	// 获取IP层信息
	ipLayer := packet.Layer(layers.LayerTypeIPv4)
	if ipLayer == nil {
		return
	}
	ip, _ := ipLayer.(*layers.IPv4)

	// 处理DNS查询
	for _, question := range dns.Questions {
		name := string(question.Name)
		var qType string

		// 确定DNS查询类型
		switch question.Type {
		case layers.DNSTypeA:
			qType = "A"
		case layers.DNSTypeAAAA:
			qType = "AAAA"
		case layers.DNSTypeCNAME:
			qType = "CNAME"
		case layers.DNSTypeMX:
			qType = "MX"
		case layers.DNSTypeNS:
			qType = "NS"
		case layers.DNSTypeTXT:
			qType = "TXT"
		default:
			qType = "OTHER"
		}

		// 创建DNS信息结构体
		dnsInfo := types.NewDnsTypes(
			name,
			qType,
			ip.SrcIP.String(),
			ip.DstIP.String(),
		)

		// 发送DNS信息到服务器
		err := sendDnsInfo(dnsInfo)
		if err != nil {
			// 处理错误，可以记录日志
		}
	}
}

// 发送DNS信息到服务器
func sendDnsInfo(dnsInfo *types.DNSTypes) error {
	data, err := json.Marshal(dnsInfo)
	if err != nil {
		return err
	}

	resp, err := http.Post("http://127.0.0.1:8080/api/v1/wsdmagent/dns", "application/json", bytes.NewBuffer(data))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return nil
}

func checkPacket(baseUrl string, connectionInfo *types.TCPConnectionInfo) bool {
	parse, err := url.Parse(baseUrl)
	if err != nil {
		return false
	}
	host := parse.Host
	ip := strings.Split(host, ":")[0]
	wsdIp := ""
	if connectionInfo.SourceIP == wsdIp && connectionInfo.DestIP == ip {
		return true
	}
	return false
}
