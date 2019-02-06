// Package agent provides definition of agent server for run tasks
package agent

// Config defines configuration for Agent
type Config struct {
	Host string
	Port string
	Dir  string
}

// Agent defines agent-server
type Agent struct {
	conf   *Config
	Master string
	events chan struct{}
}
