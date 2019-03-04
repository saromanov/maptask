package agent

import (
	"net"
	"google.golang.org/grpc"
)

// serverGRPC provides serving of the grpc server
func (aa *Agent) serveGRPC(listener net.Listener) {
	grpcServer := grpc.NewServer()
	pb.RegisterMapTaskServer(grpcServer, as)
	grpcServer.Serve(listener)
}