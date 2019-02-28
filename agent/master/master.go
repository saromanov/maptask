// Package master defines running of the master server
package master

import (
	"log"
	"net"
	"net/http"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"github.com/gorilla/mux"
)

// Run provides running of the master
func Run(address string) {
	listener, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("unable to create master server: %v", err)
	}
	defer listener.Close()

	grpcServer := grpc.NewServer()
	pb.RegisterMasterServer(grpcServer, masterServer)
	reflection.Register(grpcS)

	r := mux.NewRouter()
	r.HandleFunc("/", masterServer.uiStatusHandler)
	r.HandleFunc("/job/{id:[0-9]+}", masterServer.jobStatusHandler)
	server := &http.Server{Handler: r}

	go grpcServer.Serve(grpcL)
	go server.Serve(httpL)

	if err := m.Serve(); err != nil {
		log.Fatalf("master server failed to serve: %v", err)
	}
}