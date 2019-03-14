package maptask

import "time"

// Shard provides sharding of the input
type Shard struct {
	ID        int
	StartTime time.Time
	Dataset   *Dataset
}

// newShard creates a new input shard
func newShard(ds *Dataset) *Shard {
	return &Shard{
		Dataset:   ds,
		StartTime: time.Now().UTC(),
	}
}
