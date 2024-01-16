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

	_, err = trainClient.PurchaseTicket("London", "France", "Yatin", "Tripathi", "yatint@gmail.com", "A", 20.0)
	if err != nil {
		log.Fatalf("Failed to purchase ticket: %v", err)
	}

	//----------------------------------------------------------------------------------------------------------------

	// // Call the GetUsersBySection method
	// section := "A"
	// tickets, err := trainClient.GetUsersBySection(section)
	// if err != nil {
	// 	log.Fatalf("Failed to get users by section: %v", err)
	// }

	// // Print the received tickets
	// log.Printf("Tickets for section %s: %v", section, tickets)

	//----------------------------------------------------------------------------------------------------------------

	// // Call the GetReceiptDetails method
	// user := &pb.User{
	// 	FirstName: "Manish",
	// 	LastName:  "Singh",
	// 	Email:     "Mani@gmail.com",
	// }

	// receipt, err := trainClient.GetReceiptDetails(user.FirstName, user.LastName, user.Email)
	// if err != nil {
	// 	log.Fatalf("Failed to get receipt details: %v", err)
	// }

	// // Print the received receipt
	// log.Printf("Receipt Details: %+v", receipt)

	//----------------------------------------------------------------------------------------------------------------

	// // Call the DeleteUser method
	// userToDelete := &pb.User{
	// 	FirstName: "Manish",
	// 	LastName:  "Singh",
	// 	Email:     "Mani@gmail.com",
	// }

	// deletedTicket, err := trainClient.DeleteUser(userToDelete.FirstName, userToDelete.LastName, userToDelete.Email)
	// if err != nil {
	// 	log.Fatalf("Failed to delete user: %v", err)
	// }

	// // Print the deleted ticket
	// log.Printf("Deleted Ticket: %+v", deletedTicket)

	//----------------------------------------------------------------------------------------------------------------

	// // Call the ModifyUserSeat method
	// modifyUser := &pb.User{
	// 	FirstName: "Anurag",
	// 	LastName:  "Bansal",
	// 	Email:     "anu798@gmail.com",
	// }

	// newSeat := "B-3"
	// modifiedTicket, err := trainClient.ModifyUserSeat(modifyUser.FirstName, modifyUser.LastName, modifyUser.Email, newSeat)
	// if err != nil {
	// 	log.Fatalf("Failed to modify user seat: %v", err)
	// }

	// // Print the modified ticket
	// log.Printf("Modified Ticket: %+v", modifiedTicket)
}
