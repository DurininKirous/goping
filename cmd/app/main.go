package main

import (
	"flag"
	"fmt"
	"goping/internal/flags"
	"goping/internal/ping"
	"os"
	"time"
)

func main() {
	flags.InitFlags()

	var sent, received int
	var rtts []time.Duration

	if len(flag.Args()) != 1 {
		fmt.Println("Usage: goping [options] <address>")
		os.Exit(1)
	}

	for i := 0; i != flags.Count; i++ {
		sent++
		start := time.Now()
		ok := ping.PingOnce("0.0.0.0", "HUI", flag.Args()[0])
		if ok {
			received++
			rtts = append(rtts, time.Since(start))
		}
		time.Sleep(time.Duration(flags.Interval * float64(time.Second)))
	}

	loss := float64(sent-received) / float64(sent) * 100
	fmt.Printf("\n--- %s ping statistics ---\n", flag.Args()[0])
	fmt.Printf("%d packets transmitted, %d received, %.1f%% packet loss\n", sent, received, loss)

}
