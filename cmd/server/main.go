package main

import (
	"log"
	"net"

	"github.com/yatintri/GoAndGrpc/internal/server"
	pb "github.com/yatintri/GoAndGrpc/proto"
	"google.golang.org/grpc"
)

const (
	port = ":8080"
)

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Failed to start the server %v", err)
	}

	grpcServer := grpc.NewServer()
	// Create an instance of the server from the server package
	trainServer := server.NewTrainServer()
	pb.RegisterTrainServiceServer(grpcServer, trainServer)
	log.Printf("Server Started at %v", lis.Addr())
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to start: %v", err)
	}
}
