module github.com/birros/go-libp2p-grpc/examples/echo

go 1.16

require (
	github.com/birros/go-libp2p-grpc v0.0.0
	github.com/libp2p/go-libp2p v0.24.0
	google.golang.org/grpc v1.51.0
	google.golang.org/grpc/cmd/protoc-gen-go-grpc v1.2.0
	google.golang.org/protobuf v1.28.1
)

replace github.com/birros/go-libp2p-grpc v0.0.0 => ../..
