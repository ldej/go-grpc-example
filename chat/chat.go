package chat

import (
	"context"
	"fmt"
	"log"
)

type Server struct {
}

func (s *Server) SayHello(ctx context.Context, message *Message) (*Message, error) {
	log.Printf("Received message: %s", message.Body)
	return &Message{Body: fmt.Sprintf("Hello %s", message.Body)}, nil
}