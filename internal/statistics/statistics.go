package statistics

import "time"

func RttsStatistics(rtts []time.Duration) (time.Duration, time.Duration, time.Duration) {
	if len(rtts) == 0 {
		return 0, 0, 0
	}
	minRtts := rtts[0]
	maxRtts := rtts[0]
	var sum time.Duration
	for _, val := range rtts {
		if val < minRtts {
			minRtts = val
		}
		if val > maxRtts {
			maxRtts = val
		}
		sum += val
	}
	avgRtts := sum / time.Duration(len(rtts))
	return minRtts, avgRtts, maxRtts
}
