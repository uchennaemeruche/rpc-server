package chat

import (
	"context"

	pb "./projpb"
)

type Server struct{}

func (s *Server) SayHello(ctx context.Context, message *pb.Message) (*pb.Response, error) {

}
