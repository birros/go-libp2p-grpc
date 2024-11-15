package main

import (
	"context"
	"log"

	p2pgrpc "github.com/birros/go-libp2p-grpc"
	greeterv1 "github.com/birros/go-libp2p-grpc/examples/echo/gen/greeter/v1"
	"github.com/birros/go-libp2p-grpc/examples/echo/greeter"
	"github.com/libp2p/go-libp2p"
	"github.com/libp2p/go-libp2p/core/host"
	"github.com/libp2p/go-libp2p/core/peer"
	"github.com/libp2p/go-libp2p/core/protocol"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const pid protocol.ID = "/grpc/1.0.0"

func setupHosts(ctx context.Context) (host.Host, host.Host) {
	// hosts
	ha, _ := libp2p.New(libp2p.ListenAddrStrings("/ip4/127.0.0.1/tcp/0"))
	hb, _ := libp2p.New(libp2p.NoListenAddrs)

	// connect
	hb.Connect(ctx, peer.AddrInfo{
		ID:    ha.ID(),
		Addrs: ha.Addrs(),
	})

	return ha, hb
}

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	// ctx
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// hosts
	hs, hc := setupHosts(ctx)
	defer hs.Close()
	defer hc.Close()

	// server
	{
		// grpc server & register greeter
		s := grpc.NewServer(p2pgrpc.WithP2PCredentials())
		greeterv1.RegisterGreeterServiceServer(s, &greeter.Server{})

		// serve grpc server over libp2p host
		l := p2pgrpc.NewListener(ctx, hs, pid)
		go s.Serve(l)
	}

	// client
	{
		// client conn
		conn, _ := grpc.Dial(
			hs.ID().String(),
			grpc.WithTransportCredentials(insecure.NewCredentials()),
			p2pgrpc.WithP2PDialer(hc, pid),
		)
		defer conn.Close()

		// client
		c := greeterv1.NewGreeterServiceClient(conn)

		// SayHello
		res, _ := c.SayHello(ctx, &greeterv1.SayHelloRequest{
			Name: "Alice",
		})

		// print result
		log.Println(res.Message)
	}
}
