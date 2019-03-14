// Package agent provides definition of agent server for run tasks
package agent

import (
	"fmt"
	"path/filepath"
	"net"
)
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

// New creates a new agent
func New(c*Config) *Agent {
	dir, err := filepath.Abs(c.Dir)
	if err != nil {
		panic(err)
	}
	return &Agent {
		conf: c,
	}
}

// Start run agent
func (a*Agent) Start() error {
	listener, err := net.Listen("tcp", fmt.Sprintf("%v:%d", a.conf.Host, a.conf.Port))
	if err != nil {
		return fmt.Errorf("unable to start agent server: %v", err)
	}
	fmt.Printf("Agent server is started at %s %s\n", a.conf.Host, a.conf.Port)
}

func (a*Agent) writeConnection(reader io.Reader, writerName, channelName string, readerCount int) {
		messageWriter := NewBufferedMessageWriter(dsStore, 1000)
	
		for {
	
			message, err := util.ReadMessage(reader)
			if err == io.EOF {
				break
			}
			if err == nil {
				count += int64(len(message))
				messageWriter.WriteMessage(message)
				continue
			}
			log.Printf("on disk %s Failed to write to %s: %v", writerName, channelName, err)
		}
}
