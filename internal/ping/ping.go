package ping

import (
	"log"

	"goping/internal/icmp/reply"
	"goping/internal/icmp/request"

	"golang.org/x/net/ipv4"
)

func PingOnce(listenAddress string, messageSend string, sendToAddress string) {
	socket := request.SocketInit(listenAddress, nil)
	messageSendPayload := request.MessageInit(messageSend)
	messageSendPayloadBytes := request.MessageToBytes(messageSendPayload)
	start := request.SendIcmpPacketBytes(socket, messageSendPayloadBytes, sendToAddress)
	messageGetPayload, ipHeader, n, peer, rtt := reply.GetMessage(socket, start)
	switch messageGetPayload.Type {
	case ipv4.ICMPTypeEchoReply:
		log.Printf("got reflection from %v, len %d, ttl %d, time %.1f ms", peer, n, ipHeader.TTL, float64(rtt.Microseconds())/1000)
	default:
		log.Printf("got %+v; want echo reply", messageGetPayload)
	}
}
