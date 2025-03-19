package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"grpc_helloworld/protos/go_generated/helloWorld" // Replace with the actual import path

	"google.golang.org/grpc"
)

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	c := helloWorld.NewGreeterClient(conn)

	// Contact the server and print out its response.
	name := "World"
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// Call SayHello
	r, err := c.SayHello(ctx, &helloWorld.HelloRequest{Name: name})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	fmt.Printf("SayHello: %s\n", r.GetMessage())

	// Call SayHelloAgain
	r, err = c.SayHelloAgain(ctx, &helloWorld.HelloRequest{Name: name})
	if err != nil {
		log.Fatalf("could not greet again: %v", err)
	}
	fmt.Printf("SayHelloAgain: %s\n", r.GetMessage())
}
