package main

import (
	"context"
	"log"

	p2pgrpc "github.com/birros/go-libp2p-grpc"
	"github.com/birros/go-libp2p-grpc/examples/echo/greeter"
	"github.com/birros/go-libp2p-grpc/examples/echo/proto"
	"github.com/libp2p/go-libp2p"
	"github.com/libp2p/go-libp2p-core/host"
	"github.com/libp2p/go-libp2p-core/peer"
	"github.com/libp2p/go-libp2p-core/protocol"
	"google.golang.org/grpc"
)

const pid protocol.ID = "/grpc/1.0.0"

func setupHosts(ctx context.Context) (host.Host, host.Host) {
	// hosts
	ha, _ := libp2p.New(ctx, libp2p.ListenAddrStrings("/ip4/127.0.0.1/tcp/0"))
	hb, _ := libp2p.New(ctx, libp2p.NoListenAddrs)

	// connect
	hb.Connect(ctx, peer.AddrInfo{
		ID:    ha.ID(),
		Addrs: ha.Addrs(),
	})

	return ha, hb
}

func setupServiceRegistrar(
	ctx context.Context,
	h host.Host,
) grpc.ServiceRegistrar {
	// listener
	l := p2pgrpc.NewListener(ctx, h, pid)

	// serve
	s := grpc.NewServer(p2pgrpc.WithP2PCredentials())
	go s.Serve(l)

	return s
}

func setupClientConn(
	ctx context.Context,
	h host.Host,
	peerID peer.ID,
) (*grpc.ClientConn, error) {
	return grpc.Dial(
		peerID.String(),
		grpc.WithInsecure(),
		p2pgrpc.WithP2PDialer(ctx, h, pid),
	)
}

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	// ctx
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// hosts
	ha, hb := setupHosts(ctx)

	// service
	sr := setupServiceRegistrar(ctx, ha)
	proto.RegisterGreeterServer(sr, &greeter.Server{})

	// client
	conn, _ := setupClientConn(ctx, hb, ha.ID())
	defer conn.Close()
	c := proto.NewGreeterClient(conn)

	// SayHello
	res, _ := c.SayHello(ctx, &proto.HelloRequest{
		Name: "Alice",
	})

	// print result
	log.Println(res.Message)
}
