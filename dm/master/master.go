// Package master defines running of the master server
package master

import (
	"time"
	"log"
	"net"
	"net/http"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"github.com/gorilla/mux"
	"github.com/saromanov/maptask/agent"
)

// Server defines configuration for master server
type Server struct {
	startTime time.Time
	Topology *agent.Topology
}

// Start provides running of the master
func Start(address string) {
	listener, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("unable to create master server: %v", err)
	}
	defer listener.Close()

	server = newServer()
	grpcServer := grpc.NewServer()
	pb.RegisterMasterServer(grpcServer, masterServer)
	reflection.Register(grpcS)

	r := mux.NewRouter()
	r.HandleFunc("/", masterServer.uiStatusHandler)
	r.HandleFunc("/job/{id:[0-9]+}", masterServer.jobStatusHandler)
	server := &http.Server{Handler: r}

	go grpcServer.Serve(listener)
	go server.Serve(httpL)

	if err := m.Serve(); err != nil {
		log.Fatalf("master server failed to serve: %v", err)
	}
}

// newServer creates master server
func newServer()*Server {
	return &Server{
		startTime: time.Now().UTC(),
		Topology: agent.NewTopology(),
	}
}