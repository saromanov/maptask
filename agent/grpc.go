package agent

import (
	"net"
	"google.golang.org/grpc"
)
func (aa *Agent) serveGrpc(listener net.Listener) {
	grpcServer := grpc.NewServer()
	pb.RegisterAgentServer(grpcServer, as)
	grpcServer.Serve(listener)
}