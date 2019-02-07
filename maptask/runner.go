package maptask

import "context"
// Runner provides running of flow
type Runner interface {
	Do(*context.Context, *Flow)
}