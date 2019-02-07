package maptask

// Step represents flow step
type Step struct {
	Name string
	Fn   func([]interface{}) error
}
