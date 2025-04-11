package wsdmagent

import (
	"time"

	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"github.com/google/gopacket/tcpassembly"
	"github.com/spf13/cobra"
)

func NewWsdmAgentCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "wsdmagent",
		Short: "流量采集器",
		Long:  "一个go开发的流量采集器",
		RunE: func(cmd *cobra.Command, args []string) error {
			return runWsdmAgent()
		},
	}

	return cmd
}

func runWsdmAgent() error {
	// 创建数据包通道
	packets := make(chan gopacket.Packet, 1000)
	streamFactory := &httpStreamFactory{}
	streamPool := tcpassembly.NewStreamPool(streamFactory)
	assembler := tcpassembly.NewAssembler(streamPool)

	ticker := time.Tick(time.Minute)
	for {
		select {
		case packet := <-packets:
			if packet == nil {
				return nil
			}
			// 处理tcp/udp数据包和DNS数据包
			processPacket(packet)

			// 处理TCP流量重组
			if packet.NetworkLayer() == nil || packet.TransportLayer() == nil || packet.TransportLayer().LayerType() != layers.LayerTypeTCP {
				continue
			}

			tcp := packet.TransportLayer().(*layers.TCP)
			assembler.AssembleWithTimestamp(packet.NetworkLayer().NetworkFlow(), tcp, packet.Metadata().Timestamp)
		case <-ticker:
			// 定期清理过期的流
			assembler.FlushOlderThan(time.Now().Add(time.Minute * -2))
		}
	}
	return nil
}
