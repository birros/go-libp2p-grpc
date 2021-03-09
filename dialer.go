package p2pgrpc

import (
	"context"
	"errors"
	"net"
	"time"

	"github.com/libp2p/go-libp2p-core/host"
	"github.com/libp2p/go-libp2p-core/network"
	"github.com/libp2p/go-libp2p-core/peer"
	"github.com/libp2p/go-libp2p-core/protocol"
	"google.golang.org/grpc"
)

// WithP2PDialer uses a libp2p host as dialer. Use the id of the target host to
// create a connection. The dialer does not connect the current host to the
// target host, this must be checked before establishing a connection. It just
// wraps a gRPC connection in a libp2p stream.
func WithP2PDialer(
	ctx context.Context,
	h host.Host,
	pid protocol.ID,
) grpc.DialOption {
	return grpc.WithDialer(func(
		peerIDStr string,
		timeout time.Duration,
	) (net.Conn, error) {
		// peerID
		peerID, err := peer.IDB58Decode(peerIDStr)
		if err != nil {
			return nil, err
		}

		if h.Network().Connectedness(peerID) != network.Connected {
			return nil, errors.New("Not connected to peer")
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
