package maptask

const (
	// Input is input step
	Input = iota + 1
	// Mapper is map step
	Mapper
	// Reducer is reduce step
	Reducer
)

// Step represents flow step
type Step struct {
	id     string
	Name   string
	Fn     func([]interface{}) error
	Params map[string]interface{}
	Type   int
}

// ID returns id of the step
func (s *Step) ID() string {
	return s.id
}

// Run provides starting of the step execution
func (s *Step) Run() error {
	return s.Fn(nil)
}
