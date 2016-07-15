package main

import (
	"log"
	"net"

	pb "github.com/chrisccoy/dockergrpc/imagelist"
	"github.com/docker/engine-api/client"
	"github.com/docker/engine-api/types"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

const (
	port = ":50051"
)

// server is used to implement helloworld.GreeterServer.
type server struct{}

// SayHello implements helloworld.GreeterServer
func (s *server) ListImages(ctx context.Context, in *pb.Empty) (*pb.ImageListReply, error) {
	dCli, _ := client.NewEnvClient()
	opts := types.ImageListOptions{}
	images, _ := dCli.ImageList(context.Background(), opts)
	imgstr := make([]string, len(images))
	for i := range images {
		imgstr = append(imgstr, images[i].ID)
	}

	return &pb.ImageListReply{ImageName: imgstr}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterImageListerServer(s, &server{})
	s.Serve(lis)
}
