package gapi

import (
	"fmt"
	"grpcthing/pb"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type Server struct {
	pb.UnsafeSimpleBankServer
	config     string
	store      string
	tokenMaker string
}

func RunGrpcServer() {
	grpcServer := grpc.NewServer()
	server := &Server{
		// Initialize your server properties here
		config:     "config-value",
		store:      "store-value",
		tokenMaker: "token-maker-value",
	}

	pb.RegisterSimpleBankServer(grpcServer, server)

	reflection.Register(grpcServer)

	listener, err := net.Listen("tcp", "127.0.0.1:5000")
	if err != nil {

		fmt.Println(err)
	}
	fmt.Println("starting grpc server on 5000")
	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatal("failed t start", err)
	}
}
