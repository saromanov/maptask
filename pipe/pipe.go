// Package pipe provides mapping of the user defined tasks
// into distributed execution. It
package pipe

import "sync"

// Pipe defines struct for mapping maptask iinto distributed
type Pipe struct {
	mu *sync.Mutex
}

// New provides initialization of the Pipe
func New() *Pipe {
	return &Pipe{
		mu: &sync.Mutex{},
	}
}
