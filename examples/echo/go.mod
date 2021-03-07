module github.com/birros/go-libp2p-grpc/examples/echo

go 1.16

require (
	github.com/birros/go-libp2p-grpc v0.0.0
	github.com/golang/protobuf v1.4.2
	github.com/libp2p/go-libp2p v0.13.0
	github.com/libp2p/go-libp2p-core v0.8.5
	google.golang.org/grpc v1.36.0
	google.golang.org/protobuf v1.25.0
)

replace github.com/birros/go-libp2p-grpc v0.0.0 => ../..
