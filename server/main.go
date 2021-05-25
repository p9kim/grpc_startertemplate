package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/p9kim/grpc_startertemplate/proto"

	"google.golang.org/grpc"
)

type Server struct{}

func (s *Server) SayHello(ctx context.Context, req *proto.Message) (*proto.Message, error) {
	log.Printf("Message received from client: %s", req.Body)
	return &proto.Message{Body: "Kenobi~!!"}, nil
}

func main() {
	fmt.Println("Starting gRPC server!!")

	listener, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatal(err)
	}

	//s := proto.Server{}

	grpcServer := grpc.NewServer()

	proto.RegisterPingServer(grpcServer, &Server{})

	if err := grpcServer.Serve(listener); err != nil {
		log.Fatal(err)
	}

}
