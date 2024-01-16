package server

import (
	"context"
	"fmt"
	pb "github.com/yatintri/GoAndGrpc/proto"
)

type trainServer struct {
	pb.UnimplementedTrainServiceServer
	tickets map[string]pb.Ticket
}

func NewTrainServer() *trainServer {
	return &trainServer{
		tickets: make(map[string]pb.Ticket),
	}
}

func (s *trainServer) PurchaseTicket(ctx context.Context, req *pb.Ticket) (*pb.Ticket, error) {

	if req == nil {
		return nil, fmt.Errorf("Invalid Ticket: nil input")
	}
	if req.User == nil {
        return nil, fmt.Errorf("Invalid User: nil input")
    }
	if req.Section != "A" && req.Section != "B" {
		return nil, fmt.Errorf("Invalid Section")
	}

	seat := fmt.Sprintf("%s-%d", req.Section, len(s.tickets)+1)

	fmt.Println(seat);
	// Create a new instance of pb.Ticket before modifying it
	newReq := *req
	newReq.Section = seat

	// Initialize the User field in the new ticket
	newReq.User = &pb.User{
		FirstName: req.User.FirstName,
		LastName:  req.User.LastName,
		Email:     req.User.Email,
	}

	// Storing the new ticket in memory
	s.tickets[seat] = newReq

	return &newReq, nil
}


func (s *trainServer) GetReceiptDetails(ctx context.Context, req *pb.User) (*pb.Ticket, error) {
	// Search for the user's ticket in memory
	for _, ticket := range s.tickets {
		if ticket.User.FirstName == req.FirstName && ticket.User.LastName == req.LastName && ticket.User.Email == req.Email {
			return &ticket, nil
		}
	}

	return nil, fmt.Errorf("user not found")
}

func (s *trainServer) GetUsersBySection(req *pb.GetUsersRequest, stream pb.TrainService_GetUsersBySectionServer) error {
	// Send tickets for the requested section to the client
	for _, ticket := range s.tickets {
		if ticket.Section == req.Section {
			if err := stream.Send(&ticket); err != nil {
				return err
			}
		}
	}

	return nil
}

func (s *trainServer) DeleteUser(ctx context.Context, req *pb.User) (*pb.Ticket, error) {
	// Search for the user's ticket in memory
	for seat, ticket := range s.tickets {
		if ticket.User.FirstName == req.FirstName && ticket.User.LastName == req.LastName && ticket.User.Email == req.Email {
			// Remove the ticket from memory
			delete(s.tickets, seat)
			return &ticket, nil
		}
	}

	return nil, fmt.Errorf("user not found")
}

func (s *trainServer) ModifyUserSeat(ctx context.Context, req *pb.ModifyUserSeatRequest) (*pb.Ticket, error) {
	// Search for the user's ticket in memory
	for seat, ticket := range s.tickets {
		if ticket.User.FirstName == req.User.FirstName && ticket.User.LastName == req.User.LastName && ticket.User.Email == req.User.Email {
			// Modify the user's seat
			oldSeat := ticket.Section
			ticket.Section = req.NewSeat

			// Update the ticket in memory
			s.tickets[seat] = ticket
			delete(s.tickets, oldSeat)

			return &ticket, nil
		}
	}

	return nil, fmt.Errorf("user not found")
}
