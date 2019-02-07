package maptask

// Flow defines the main flow for the app
type Flow struct {
}

// NewText provides initialization of the flow
// with a text
func NewText(text string) *Flow {
	return &Flow{}
}

// Map provides mapping of the function
func (f* Flow) Map(name string, fn func(interface{}) error) *Flow{
	return f
}

// Reduce
func (f* Flow) Reduce(name string, fn func(interface{}) error) *Flow{
	return f
}
