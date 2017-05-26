package main

import (
	"log"
	"net"
	"strings"

	"golang.org/x/net/context"
	"google.golang.org/grpc"

	pb "github.com/frozzare/grpc-go-php/go/user"
)

const (
	port = ":50051"
)

// server is used to implement customer.CustomerServer.
type server struct {
	users []*pb.UserRequest
}

// CreateCustomer creates a new Customer
func (s *server) CreateUser(ctx context.Context, in *pb.UserRequest) (*pb.UserResponse, error) {
	s.users = append(s.users, in)
	return &pb.UserResponse{Id: in.Id}, nil
}

// GetCustomers returns all customers by given filter
func (s *server) GetUsers(filter *pb.UserFilter, stream pb.User_GetUsersServer) error {
	for _, user := range s.users {
		if filter.Name != "" {
			if !strings.Contains(user.Name, filter.Name) {
				continue
			}
		}
		if err := stream.Send(user); err != nil {
			return err
		}
	}
	return nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	// Creates a new gRPC server
	s := grpc.NewServer()
	pb.RegisterUserServer(s, &server{})
	s.Serve(lis)
}
