// Package pipe provides mapping of the user defined tasks
// into distributed execution. It
package pipe

import (
	"sync"
	"github.com/saromanov/maptask/proto"
)

// Pipe defines struct for mapping maptask iinto distributed
type Pipe struct {
	mu *sync.Mutex
	steps []*proto.FlowStep
}

// New provides initialization of the Pipe
func New(steps []*proto.FlowStep) *Pipe {
	return &Pipe{
		mu: &sync.Mutex{},
		steps: steps,
	}
}
