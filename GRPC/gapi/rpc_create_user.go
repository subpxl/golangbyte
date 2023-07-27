package gapi

import (
	"context"
	"fmt"
	"grpcthing/pb"
)

var MyDB = DbInit()

func (s *Server) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	user := DBUser{username: req.Username, email: req.Email, password: req.Password, full_name: req.FullName}
	InsertData(MyDB, user)
	if err != nil {
		fmt.Println(err)

		// return nil, status.Errorf(codes.Internal, "user not created")
	}

	return &pb.CreateUserResponse{User: convertUser(user)}, nil
}
