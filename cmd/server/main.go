package main

import (
	"log"
	"net"

	pb "github.com/yatintri/GoAndGrpc/proto"
	"google.golang.org/grpc"
)

const (
	port = ":8080"
)

type server struct {
	pb.TrainServiceServer
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Failed to start the server %v", err)
	}
	grpcServer := grpc.NewServer()
	pb.RegisterTrainServiceServer(grpcServer, &server{})
	log.Printf("Server Started at %v", lis.Addr())
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to start: %v", err)
	}
}
