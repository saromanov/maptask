package maptask

// Step represents flow step
type Step struct {
	id     string
	Name   string
	Fn     func([]interface{}) error
	Params map[string]interface{}
}

// ID returns id of the step
func (s *Step) ID() string {
	return s.id
}

// Run provides starting of the step execution
func (s *Step) Run() error {
	return s.Fn(nil)
}
