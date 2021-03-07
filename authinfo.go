package p2pgrpc

import (
	"context"

	"github.com/libp2p/go-libp2p-core/network"
	"github.com/libp2p/go-libp2p-core/peer"
	"google.golang.org/grpc/credentials"
	grpcpeer "google.golang.org/grpc/peer"
)

var _ credentials.AuthInfo = (*AuthInfo)(nil)

// AuthInfo ...
type AuthInfo struct {
	Stream network.Stream
}

// AuthType ...
func (ai AuthInfo) AuthType() string {
	return ""
}

// AuthInfoFromContext ...
func AuthInfoFromContext(ctx context.Context) (AuthInfo, bool) {
	p, ok := grpcpeer.FromContext(ctx)
	if !ok {
		return AuthInfo{}, false
	}

	ai, ok := p.AuthInfo.(AuthInfo)
	return ai, ok
}

// RemotePeerFromContext ...
func RemotePeerFromContext(ctx context.Context) (peer.ID, bool) {
	ai, ok := AuthInfoFromContext(ctx)
	if !ok {
		return "", false
	}

	peerID := ai.Stream.Conn().RemotePeer()
	return peerID, true
}
