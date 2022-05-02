package server

import (
	"context"

	"github.com/jeffotoni/grpc-crud/proto"
)

type Server struct {
	proto.UnimplementedCustomersServiceServer
}

func New() *Server {
	return &Server{}
}

func (s *Server) Get(ctx context.Context, req *proto.GetRequest) (*proto.Customers, error) {
	return &proto.Customers{}, nil
}
