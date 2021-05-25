package main

import (
	"context"
	"log"

	"github.com/p9kim/grpc_startertemplate/proto"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:9000", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}

	defer conn.Close()

	client := proto.NewPingClient(conn)

	message := proto.Message{
		Body: "Hello There!",
	}

	res, err := client.SayHello(context.Background(), &message)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Response from server: %s", res.Body)

}
