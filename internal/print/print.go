package print

import (
	"flag"
	"fmt"
	"goping/internal/statistics"
	"time"
)

func PrintStatistics(received int, sent int, rtts []time.Duration) {
	loss := float64(sent-received) / float64(sent) * 100
	minRtts, avgRtts, maxRtts := statistics.RttsStatistics(rtts)
	fmt.Printf("\n--- %s ping statistics ---\n", flag.Args()[0])
	fmt.Printf("%d packets transmitted, %d received, %.1f%% packet loss\n", sent, received, loss)
	fmt.Printf("rtt min/avg/max = %.2f/%.2f/%.2f ms\n", float64(minRtts.Microseconds())/1000, float64(avgRtts.Microseconds())/1000, float64(maxRtts.Microseconds())/1000)
}
