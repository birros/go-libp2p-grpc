package p2pgrpc

import (
	"context"
	"errors"
	"net"

	"github.com/libp2p/go-libp2p-core/network"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

var _ credentials.TransportCredentials = (*creds)(nil)

type creds struct{}

// WithP2PCredentials ...
func WithP2PCredentials() grpc.ServerOption {
	return grpc.Creds(creds{})
}

func (f creds) ClientHandshake(ctx context.Context, _ string, c net.Conn) (net.Conn, credentials.AuthInfo, error) {
	return c, nil, nil
}

func (f creds) ServerHandshake(c net.Conn) (net.Conn, credentials.AuthInfo, error) {
	s, ok := c.(network.Stream)
	if !ok {
		return nil, nil, errors.New("Not a libp2p Stream")
	}

	i := AuthInfo{
		Stream: s,
	}
	return c, i, nil
}

func (f creds) Info() credentials.ProtocolInfo {
	return credentials.ProtocolInfo{
		ProtocolVersion:  "",
		SecurityProtocol: "",
		SecurityVersion:  "",
		ServerName:       "",
	}
}

func (f creds) Clone() credentials.TransportCredentials {
	return creds{}
}

func (f creds) OverrideServerName(string) error {
	return nil
}
