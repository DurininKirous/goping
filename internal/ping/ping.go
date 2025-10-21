package ping

import (
	"fmt"
	"log"
	"os"
	"time"

	"goping/internal/icmp/reply"
	"goping/internal/icmp/request"

	"golang.org/x/net/ipv4"
)

func PingOnce(listenAddress string, messageSend string, sendToAddress string, timeout float64) bool {
	opts := &request.PingOptions{ProtocolName: "ip4:icmp", TimeoutSeconds: timeout}
	socket := request.SocketInit(listenAddress, opts)

	messageSendPayload := request.MessageInit(messageSend)
	messageSendPayloadBytes := request.MessageToBytes(messageSendPayload)
	start := request.SendIcmpPacketBytes(socket, messageSendPayloadBytes, sendToAddress)
	deadline := start.Add(time.Duration(timeout * float64(time.Second)))
	socket.SetReadDeadline(deadline)

	messageGetPayload, ipHeader, n, peer, rtt, err := reply.GetMessage(socket, start)
	if err != nil {
		if os.IsTimeout(err) {
			fmt.Println("Request timed out")
		} else {
			log.Fatal(err)
		}
		socket.Close()
		return false
	}
	switch messageGetPayload.Type {
	case ipv4.ICMPTypeEchoReply:
		log.Printf("got reflection from %v, len %d, ttl %d, time %.1f ms", peer, n, ipHeader.TTL, float64(rtt.Microseconds())/1000)
		return true
	default:
		log.Printf("got %+v; want echo reply", messageGetPayload)
		return false
	}
}
