package agent

import "sync"

// AgentInfo provides definition for agent information
type AgentInfo struct {
}

// Stand contains group of Agents
type Stand struct {
	sync.RWMutex
	Name   string
	Agents map[string]*AgentInfo
}

// DataCenter contains group of Stands
type DataCenter struct {
	sync.RWMutex
	Name   string
	Stands map[string]*Stand
}

// Topology defines representation of datacenters
type Topology struct {
	mu          *sync.RWMutex
	DataCenters map[string]*DataCenter
}

// NewTopology creates topology instance
func NewTopology() *Topology {
	return &Topology{
		mu:          &sync.Mutex{},
		DataCenters: map[string]*DataCenter{},
	}
}
