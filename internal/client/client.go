package client

import (
	"context"
	"fmt"

	pb "github.com/yatintri/GoAndGrpc/proto"
)

type TrainClient struct {
	Client pb.TrainServiceClient
}

func NewTrainClient(client pb.TrainServiceClient) *TrainClient {
	return &TrainClient{
		Client: client,
	}
}

func (tc *TrainClient) PurchaseTicket(from, to, firstName, lastName, email, section string, price float32) (*pb.Ticket, error) {
	user := &pb.User{
		FirstName: firstName,
		LastName:  lastName,
		Email:     email,
	}

	ticket := &pb.Ticket{
		From:    from,
		To:      to,
		User:    user,
		Price:   price,
		Section: section,
	}

	response, err := tc.Client.PurchaseTicket(context.Background(), ticket)

	if err != nil {
		return nil, fmt.Errorf("PurchaseTicket failed: %v", err)
	}

	return response, nil
}

func (tc *TrainClient) GetReceiptDetails(firstName, lastName, email string) (*pb.Ticket, error) {
	user := &pb.User{
		FirstName: firstName,
		LastName:  lastName,
		Email:     email,
	}

	response, err := tc.Client.GetReceiptDetails(context.Background(), user)
	if err != nil {
		return nil, fmt.Errorf("GetReceiptDetails failed: %v", err)
	}

	return response, nil
}

func (tc *TrainClient) GetUsersBySection(section string) ([]*pb.Ticket, error) {
	request := &pb.GetUsersRequest{
		Section: section,
	}

	stream, err := tc.Client.GetUsersBySection(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("GetUsersBySection failed: %v", err)
	}

	var tickets []*pb.Ticket

	for {
		ticket, err := stream.Recv()
		if err != nil {
			break
		}
		tickets = append(tickets, ticket)
	}

	return tickets, nil
}

func (tc *TrainClient) DeleteUser(firstName, lastName, email string) (*pb.Ticket, error) {
	user := &pb.User{
		FirstName: firstName,
		LastName:  lastName,
		Email:     email,
	}

	response, err := tc.Client.DeleteUser(context.Background(), user)
	if err != nil {
		return nil, fmt.Errorf("DeleteUser failed: %v", err)
	}

	return response, nil
}

func (tc *TrainClient) ModifyUserSeat(firstName, lastName, email, newSeat string) (*pb.Ticket, error) {
	user := &pb.User{
		FirstName: firstName,
		LastName:  lastName,
		Email:     email,
	}

	request := &pb.ModifyUserSeatRequest{
		User:    user,
		NewSeat: newSeat,
	}

	response, err := tc.Client.ModifyUserSeat(context.Background(), request)
	if err != nil {
		return nil, fmt.Errorf("ModifyUserSeat failed: %v", err)
	}

	return response, nil
}