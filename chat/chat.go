package chat

import (
	"context"
	"fmt"
	"log"
)

type Server struct{}

func (s *Server) SayHello(ctx context.Context, message *Message) (*Response, error) {
	log.Println("SayHello called")
	return &Response{Body: fmt.Sprintf("New message: %s", message.Body)}, nil
}

func (s *Server) GetDetails(ctx context.Context, message *Details) (*Response, error) {
	log.Println("GetDetails called")
	return &Response{Body: fmt.Sprintf("Your name is %s and you're %d years old", message.Name, message.Age)}, nil
}

func (s *Server) mustEmbedUnimplementedChatServiceServer() {}
