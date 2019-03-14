package maptask

import "time"

// Shard provides sharding of the input
type Shard struct {
	ID        int
	StartTime time.Time
	Dataset   *Dataset
}
