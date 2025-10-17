package main

import (
	"goping/internal/flags"
	"goping/internal/ping"
	"time"
)

func main() {
	flags.InitFlags()
	for i := 0; i != flags.Count; i++ {
		ping.PingOnce("0.0.0.0", "HUI", "8.8.8.8")
		time.Sleep(time.Duration(flags.Interval * float64(time.Second)))
	}
}
