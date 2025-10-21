package request

import (
	"bytes"
	"log"
	"net"
	"os"
	"time"

	"golang.org/x/net/icmp"
	"golang.org/x/net/ipv4"
)

const DefaultPayloadSize = 56

type PingOptions struct {
	ProtocolName   string
	TimeoutSeconds float64
}

func SocketInit(listenAddress string, opts *PingOptions) *icmp.PacketConn {
	if opts == nil {
		opts = &PingOptions{ProtocolName: "ip4:icmp", TimeoutSeconds: 2}
	}
	c, err := icmp.ListenPacket(opts.ProtocolName, listenAddress)
	if err != nil {
		log.Fatal(err)
	}

	return c
}

func MessageInit(message string) icmp.Message {
	data := []byte(message)

	if len(data) < DefaultPayloadSize {
		padding := bytes.Repeat([]byte("Q"), DefaultPayloadSize-len(data))
		data = append(data, padding...)
	} else if len(data) > DefaultPayloadSize {
		data = data[:DefaultPayloadSize]
	}

	wm := icmp.Message{
		Type: ipv4.ICMPTypeEcho, Code: 0,
		Body: &icmp.Echo{
			ID: os.Getpid() & 0xffff, Seq: 1,
			Data: data,
		},
	}

	return wm
}

func MessageToBytes(message icmp.Message) []byte {
	wb, err := message.Marshal(nil)
	if err != nil {
		log.Fatal(err)
	}

	return wb
}

func SendIcmpPacketBytes(socket *icmp.PacketConn, message []byte, sendAdress string) *time.Time {
	start := time.Now()
	if _, err := socket.WriteTo(message, &net.IPAddr{IP: net.ParseIP(sendAdress)}); err != nil {
		log.Fatal(err)
	}

	return &start
}
