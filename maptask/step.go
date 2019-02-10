package maptask

// Step represents flow step
type Step struct {
	ID     string
	Name   string
	Fn     func([]interface{}) error
	Params map[string]interface{}
}

// Run provides starting of the step execution
func (s *Step) Run() error {
	return s.Fn(nil)
}
