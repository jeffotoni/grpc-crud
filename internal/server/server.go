package server

import (
	"context"

	"github.com/jeffotoni/grpc-crud/proto"
)

type Server struct {
	proto.UnimplementedPoolServiceServer
}

func New() *Server {
	return &Server{}
}

func (s *Server) Get(ctx context.Context, req *proto.GetRequest) (*proto.Customers, error) {
	return &proto.Customers{}, nil
}
