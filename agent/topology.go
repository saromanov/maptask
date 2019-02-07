package agent

import "sync"

// AgentInfo provides definition for agent information
type AgentInfo struct {
}

// Stand contains group of Agents
type Stand struct {
	mu     *sync.RWMutex
	Name   string
	Agents map[string]*AgentInfo
}

// DataCenter contains group of Stands
type DataCenter struct {
	mu     *sync.RWMutex
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
		DataCenters: make(map[string]*DataCenter),
	}
}

// NewDataCenter create s anew data center
func NewDataCenter(name string) *DataCenter {
	return &DataCenter{
		mu:          &sync.Mutex{},
		Name :name,
		Stands: make(map[string]*Stand)m

	}
}
