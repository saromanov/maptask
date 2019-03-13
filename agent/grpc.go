package agent

import (
	"net"
	"google.golang.org/grpc"
	"github.com/saromanov/maptask/proto"
)

// serverGRPC provides serving of the grpc server
func (aa *Agent) serveGRPC(listener net.Listener) {
	grpcServer := grpc.NewServer()
	proto.RegisterMapTaskServer(grpcServer, as)
	grpcServer.Serve(listener)
}