package main

import (
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/yatintri/GoAndGrpc/internal/client"
	pb "github.com/yatintri/GoAndGrpc/proto"
)

const (
	port = ":8080"
)

func main() {
	conn, err := grpc.Dial("localhost"+port, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Did not connect: %v", err)
	}
	defer conn.Close()

	trainClient := client.NewTrainClient(pb.NewTrainServiceClient(conn))

	// Use the underscore to ignore the error if you're not planning to handle it explicitly
	_, err = trainClient.PurchaseTicket("London", "France", "Yatin", "Tripathi", "yatintripathi0679@gmail.com", "A", 20.0)
	if err != nil {
		log.Fatalf("Failed to purchase ticket: %v", err)
	}

}
