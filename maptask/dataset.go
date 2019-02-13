package maptask

// Dataset implements representation of the tasks
type Dataset struct {
	Flow *Flow
}

// newDataset provides initialization of the dataset
func newDataset(f *Flow) *Dataset {
	d := &Dataset{
		Flow: f,
	}
	return d
}
