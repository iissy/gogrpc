package main

import (
	"log"
	"net"

	"github.com/iissy/gogrpc/helloworld"
	"github.com/iissy/gogrpc/messages"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	port = ":50052"
)

type server struct{}

func (s *server) SayHello(ctx context.Context, in *messages.HelloRequest) (*messages.HelloReply, error) {
	return &messages.HelloReply{Message: "Hello " + in.Name}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	helloworld.RegisterGreeterServer(s, &server{})
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
