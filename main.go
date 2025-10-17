package main

import (
	"log"

	"goping/internal/icmp/reply"
	"goping/internal/icmp/request"

	"golang.org/x/net/ipv4"
)

func main() {
	socket := request.SocketInit("0.0.0.0", nil)
	message := request.MessageInit("Hello")
	messageBytes := request.MessageToBytes(message)
	start := request.SendIcmpPacketBytes(socket, messageBytes, "192.168.1.1")
	messageGet, ipHeader, n, peer, rtt := reply.GetMessage(socket, start)
	switch messageGet.Type {
	case ipv4.ICMPTypeEchoReply:
		log.Printf("got reflection from %v, len %d, ttl %d, time %.1f ms", peer, n, ipHeader.TTL, float64(rtt.Microseconds())/1000)
	default:
		log.Printf("got %+v; want echo reply", messageGet)
	}
}
