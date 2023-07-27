package gapi

import "grpcthing/pb"

func convertUser(user DBUser) *pb.User {
	return &pb.User{
		Username: user.username,
		FullName: user.full_name,
		Email:    user.email,
	}
}
