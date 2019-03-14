// Package pipe provides mapping of the user defined tasks
// into distributed execution. It
package pipe

import (
	"errors"
	"sync"
	pb "github.com/saromanov/maptask/proto"
)

var errNoSteps = errors.New("steps is not defined")

// Pipe defines struct for mapping maptask iinto distributed
type Pipe struct {
	mu *sync.Mutex
	steps []*pb.FlowStep
}

// New provides initialization of the Pipe
func New(steps []*pb.FlowStep) (*Pipe, error) {
	if len(steps) == 0 {
		return nil, errNoSteps
	}
	return &Pipe{
		mu: &sync.Mutex{},
		steps: steps,
	}, nil
}
