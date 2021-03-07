package greeter

import (
	"context"
	"errors"
	"log"

	p2pgrpc "github.com/birros/go-libp2p-grpc"
	"github.com/birros/go-libp2p-grpc/examples/echo/proto"
)

var _ proto.GreeterServer = (*Server)(nil)

// Server ...
type Server struct{}

// SayHello ...
func (s *Server) SayHello(
	ctx context.Context,
	req *proto.HelloRequest,
) (*proto.HelloResponse, error) {
	peerID, ok := p2pgrpc.RemotePeerFromContext(ctx)
	if !ok {
		return nil, errors.New("No AuthInfo in context")
	}

	log.Println("Request from: " + peerID.String())

	res := &proto.HelloResponse{
		Message: "Hello " + req.Name + "!",
	}
	return res, nil
}
