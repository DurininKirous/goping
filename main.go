package main

import (
	"log"
	"net"
	"os"
	"time"

	"golang.org/x/net/icmp"
	"golang.org/x/net/ipv4"
)

func main() {
	c, err := icmp.ListenPacket("ip4:icmp", "0.0.0.0")
	if err != nil {
		log.Fatal(err)
	}
	defer c.Close()

	wm := icmp.Message{
		Type: ipv4.ICMPTypeEcho, Code: 0,
		Body: &icmp.Echo{
			ID: os.Getpid() & 0xffff, Seq: 1,
			Data: []byte("HELLO-R-U-THERE"),
		},
	}

	wb, err := wm.Marshal(nil)
	if err != nil {
		log.Fatal(err)
	}

	start := time.Now()
	if _, err := c.WriteTo(wb, &net.IPAddr{IP: net.ParseIP("8.8.8.8")}); err != nil {
		log.Fatal(err)
	}

	rb := make([]byte, 1500)
	n, peer, err := c.ReadFrom(rb)
	if err != nil {
		log.Fatal(err)
	}
	rtt := time.Since(start)

	rm, err := icmp.ParseMessage(1, rb[:n])
	if err != nil {
		log.Fatal(err)
	}
	iph, err := ipv4.ParseHeader(rb)
	switch rm.Type {
	case ipv4.ICMPTypeEchoReply:
		log.Printf("got reflection from %v, len %d, ttl %d, time %.1f ms", peer, n, iph.TTL, float64(rtt.Microseconds())/1000)
	default:
		log.Printf("got %+v; want echo reply", rm)
	}
}
