package main

import (
	"flag"
	"fmt"
	"goping/internal/flags"
	"goping/internal/ping"
	"goping/internal/print"
	"os"
	"os/signal"
	"syscall"
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
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)

	go func() {
		<-sigChan
		fmt.Println("\n\n--- Interrupted ---")
		print.PrintStatistics(received, sent, rtts)
		os.Exit(0)
	}()

	for i := 0; i != flags.Count; i++ {
		sent++
		start := time.Now()
		ok := ping.PingOnce("0.0.0.0", "HUI", flag.Args()[0], flags.Timeout)
		if ok {
			received++
			rtts = append(rtts, time.Since(start))
		}
		time.Sleep(time.Duration(flags.Interval * float64(time.Second)))
	}
	print.PrintStatistics(received, sent, rtts)

}
