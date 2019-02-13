package maptask

import "context"

// Flow defines the main flow for the app
type Flow struct {
	steps []*Step
	name  string
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

// Map provides mapping of the function
func (f *Flow) Map(name string, fn func([]interface{}) error) *Flow {
	step := &Step{
		Name: name,
		Fn:   fn,
	}
	f.steps = append(f.steps, step)
	return f
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
	}
	f.steps = append(f.steps, step)
	return f
}
