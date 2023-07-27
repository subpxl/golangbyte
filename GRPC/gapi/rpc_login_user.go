package gapi

import (
	"context"
	"grpcthing/pb"
)

func (s *Server) LoginUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	user := LoginData(DB, DBUser{username: req.Username, password: req.Password})
	return &pb.CreateUserResponse{User: convertUser(*user)}, nil

}
