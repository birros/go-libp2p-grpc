all: run

.PHONY: run
run:
	go run main/main.go

.PHONY: setup-protobuf
setup-protobuf:
	go install \
		google.golang.org/protobuf/cmd/protoc-gen-go \
		google.golang.org/grpc/cmd/protoc-gen-go-grpc
	hash asdf > /dev/null 2>&1 && \
		asdf reshim golang

.PHONY: generate
generate: setup-protobuf
	buf generate
