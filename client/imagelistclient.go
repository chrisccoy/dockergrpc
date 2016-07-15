package main

import (
	"log"

	pb "github.com/chrisccoy/dockergrpc/imagelist"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

const (
	address = "localhost:50051"
)

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewImageListerClient(conn)

	// Contact the server and print out its response.
	r, err := c.ListImages(context.Background(), &pb.Empty{})
	if err != nil {
		log.Fatalf("could not list: %v", err)
	}
	log.Printf("Listing: %s", r.ImageName)
}
