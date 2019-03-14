package maptask

import (
	"context"
	"fmt"
	"os"
)

// Flow defines the main flow for the app
type Flow struct {
	steps  []*Step
	name   string
	shards int
}

// NewText provides initialization of the flow
// with a text
func NewText(text string) *Flow {
	return &Flow{
		steps: []*Step{},
	}
}

// New provides initialization of the flow
// with name
func New(name string) *Flow {
	return &Flow{
		steps: []*Step{},
		name:  name,
	}
}

// Read provides reading of the data from file
func (f *Flow) Read(path string) error {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return fmt.Errorf("file is not exist: %v", err)
	}
	f.steps = append(f.steps, &Step{
		Name: "input",
		Type: Input,
	})
	return nil
}

// Map provides mapping of the function
func (f *Flow) Map(name string, fn func([]interface{}) error) *Flow {
	step := &Step{
		Name: name,
		Fn:   fn,
		Type: Mapper,
	}
	f.steps = append(f.steps, step)
	return f
}

// AddShardsCount provides adding of the shards count to the flow
func (f *Flow) AddShardsCount(num int) {
	f.shards = num
}

// Run provides excecuting of the flow
func (f *Flow) Run(c *context.Context) error {
	return f.run(c)
}

func (f *Flow) run(c *context.Context) error {

	return nil
}

// Reduce
func (f *Flow) Reduce(name string, fn func([]interface{}) error) *Flow {
	step := &Step{
		Name: name,
		Fn:   fn,
		Type: Reducer,
	}
	f.steps = append(f.steps, step)
	return f
}
