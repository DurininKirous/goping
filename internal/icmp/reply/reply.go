package reply

import (
	"net"
	"time"

	"golang.org/x/net/icmp"
	"golang.org/x/net/ipv4"
)

const DefaultMTU = 1500

func GetMessageBytes(socket *icmp.PacketConn, start *time.Time) ([]byte, int, net.Addr, time.Duration, error) {
	readBytes := make([]byte, DefaultMTU)
	n, peer, err := socket.ReadFrom(readBytes)
	if err != nil {
		return nil, 0, nil, 0, err
	}
	rtt := time.Since(*start)

	return readBytes, n, peer, rtt, err
}

func GetMessage(socket *icmp.PacketConn, start *time.Time) (*icmp.Message, *ipv4.Header, int, net.Addr, time.Duration, error) {
	readBytes, n, peer, rtt, err := GetMessageBytes(socket, start)
	if err != nil {
		return nil, nil, 0, nil, 0, err
	}

	var readOffSet int
	ipHeader, err := ipv4.ParseHeader(readBytes)
	if err != nil {
		return nil, nil, 0, nil, 0, err
	}
	if ipHeader.Len < n {
		readOffSet = ipHeader.Len
	} else {
		readOffSet = 0
	}

	readMessage, err := icmp.ParseMessage(1, readBytes[readOffSet:n])
	if err != nil {
		return nil, nil, 0, nil, 0, err
	}

	return readMessage, ipHeader, n, peer, rtt, nil
}
