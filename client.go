package main

import (
	"context"
	"log"

	"google.golang.org/grpc"

	"github.com/ldej/go-chat/chat"
)

func main() {
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":9000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to dial at port 9000: %v", err)
	}
	defer conn.Close()

	c := chat.NewChatServiceClient(conn)
	response, err := c.SayHello(context.Background(), &chat.Message{
		Body: "Laurence",
	})
	if err != nil {
		log.Fatalf("Error %v", err)
	}
	log.Printf("Reponse: %s", response.Body)
}