package main

import (
	"github.com/pavel-one/EdgeGPT-Go"
	"github.com/pavel-one/EdgeGPT-Go/internal/GRPC"
	"github.com/pavel-one/EdgeGPT-Go/internal/Logger"
	pb "github.com/pavel-one/EdgeGPT-Go/pkg/GRPC/GPT"
	"google.golang.org/grpc"
	"net"
)

var log = Logger.NewLogger("General")
var storage = EdgeGPT.NewStorage()

func main() {
	srv := GRPC.NewServer(storage)

	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterGptServiceServer(s, srv)
	log.Infoln("Starting server on port 8080")

	if err := s.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
