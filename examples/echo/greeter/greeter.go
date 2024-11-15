package greeter

import (
	"context"
	"errors"
	"log"

	p2pgrpc "github.com/birros/go-libp2p-grpc"
	greeterv1 "github.com/birros/go-libp2p-grpc/examples/echo/gen/greeter/v1"
)

var _ greeterv1.GreeterServiceServer = (*Server)(nil)

// Server ...
type Server struct{}

// SayHello ...
func (s *Server) SayHello(
	ctx context.Context,
	req *greeterv1.SayHelloRequest,
) (*greeterv1.SayHelloResponse, error) {
	peerID, ok := p2pgrpc.RemotePeerFromContext(ctx)
	if !ok {
		return nil, errors.New("no AuthInfo in context")
	}

	log.Println("Request from: " + peerID.String())

	res := &greeterv1.SayHelloResponse{
		Message: "Hello " + req.Name + "!",
	}
	return res, nil
}
