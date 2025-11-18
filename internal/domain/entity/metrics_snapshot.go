package entity

import "time"

type MetricsSnapshot struct {
	Hits            uint64
	Misses            uint64
	UpstreamLatency time.Duration
	CacheLatency	time.Duration
	Evictions	  uint64
}