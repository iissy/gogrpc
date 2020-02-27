package main

import (
	"log"
	"os"

	"github.com/iissy/gogrpc/helloworld"
	"github.com/iissy/gogrpc/messages"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

const (
	address     = "0.0.0.0:50052"
	defaultName = "world"
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := helloworld.NewGreeterClient(conn)

	name := defaultName
	if len(os.Args) > 1 {
		name = os.Args[1]
	}
	r, err := c.SayHello(context.Background(), &messages.HelloRequest{Name: name})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r.Message)
}
