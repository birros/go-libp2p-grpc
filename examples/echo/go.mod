module github.com/birros/go-libp2p-grpc/examples/echo

go 1.16

require (
	github.com/birros/go-libp2p-grpc v0.0.0
	github.com/libp2p/go-libp2p v0.15.0
	github.com/libp2p/go-libp2p-core v0.9.0
	google.golang.org/grpc v1.40.0
	google.golang.org/protobuf v1.27.1
)

replace github.com/birros/go-libp2p-grpc v0.0.0 => ../..
