package maptask

import "time"

// Dataset implements representation of the tasks
type Dataset struct {
	Flow      *Flow
	StartTime time.Time
	Shards    []*Shard
}

// newDataset provides initialization of the dataset
func newDataset(f *Flow) *Dataset {
	d := &Dataset{
		Flow:      f,
		StartTime: time.Now().UTC(),
	}
	return d
}

// newSharedDataset creates dataset with shards
func newShardDataset(f *Flow, shardCount int) *Dataset {
	ds := newDataset(f)
	for i := 0; i < shardCount; i++ {
		ds.Shards = append(ds.Shards, newShard(ds))
	}
	return ds
}
