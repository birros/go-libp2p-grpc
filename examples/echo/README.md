# Echo example

Example of use of `go-libp2p-grpc`.

## Install

```shell
$ brew install protobuf
$ go get google.golang.org/protobuf/cmd/protoc-gen-go \
    google.golang.org/grpc/cmd/protoc-gen-go-grpc
$ brew tap bufbuild/buf
$ brew install buf
```

## Usage

```shell
$ make
$ make generate # protobuf go files
```
