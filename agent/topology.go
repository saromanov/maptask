package agent

import "sync"

// AgentInfo provides definition for agent information
type AgentInfo struct {
}

type Stand struct {
	sync.RWMutex
	Name   string
	Agents map[string]*AgentInfo
}

type DataCenter struct {
	sync.RWMutex
	Name   string
	Stands map[string]*Stand
}

type Topology struct {
	sync.RWMutex
	DataCenters map[string]*DataCenter
}
