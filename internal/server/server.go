package server

import (
	"context"
	"fmt"

	"github.com/jeffotoni/grpc-crud/proto"
)

type Server struct {
	proto.UnimplementedCustomersServiceServer
}

func New() *Server {
	return &Server{}
}

func (s *Server) Get(ctx context.Context, req *proto.GetRequest) (*proto.Customers, error) {

	fmt.Println("get:", req.Document)
	return &proto.Customers{
		Document: "TID-00002",
	}, nil
}
