package p2pgrpc

import (
	"context"
	"net"
	"time"

	"github.com/libp2p/go-libp2p-core/host"
	"github.com/libp2p/go-libp2p-core/peer"
	"github.com/libp2p/go-libp2p-core/protocol"
	"google.golang.org/grpc"
)

// WithP2PDialer ...
func WithP2PDialer(ctx context.Context, h host.Host, pid protocol.ID) grpc.DialOption {
	return grpc.WithDialer(func(peerIDStr string, timeout time.Duration) (net.Conn, error) {
		// peerID
		peerID, err := peer.IDB58Decode(peerIDStr)
		if err != nil {
			return nil, err
		}

		// ctx
		ctx, ctxCancel := context.WithTimeout(ctx, timeout)
		defer ctxCancel()

		// stream
		stream, err := h.NewStream(ctx, peerID, pid)
		if err != nil {
			return nil, err
		}

		return &Conn{Stream: stream}, nil
	})
}
