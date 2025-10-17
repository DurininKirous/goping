package flags

import "flag"

var (
	PingAddress string
	Count       int
	Interval    float64
)

func InitFlags() {
	flag.IntVar(&Count, "c", -1, "Number of pings")
	flag.Float64Var(&Interval, "i", 1, "Interval between pings")
	flag.Parse()
}
