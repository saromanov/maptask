package maptask

import "time"

// Dataset implements representation of the tasks
type Dataset struct {
	Flow *Flow
	StartTime time.Time
}

// newDataset provides initialization of the dataset
func newDataset(f *Flow) *Dataset {
	d := &Dataset{
		Flow: f,
		StartTime: time.Now().UTC(),
	}
	return d
}
