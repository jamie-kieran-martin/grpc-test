package chat

import (
	"fmt"
	"log"

	"golang.org/x/net/context"
)

type Server struct {
	UnimplementedChatServiceServer
}

func (s *Server) SayHello(ctx context.Context, in *Message) (*Response, error) {
	log.Printf("Received from client: %s", in.Name)
	return &Response{Result: fmt.Sprintf("Hello %v!", in.Name)}, nil
}
