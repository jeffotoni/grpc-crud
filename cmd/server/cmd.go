package main

import (
	"context"
	"log"
	"net"
	"time"

	"github.com/jeffotoni/grpc-crud/internal/server"
	pb "github.com/jeffotoni/grpc-crud/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
)

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	//
	// GRPC
	//

	lis, err := net.Listen("tcp", "0.0.0.0:4001")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	unary := []grpc.UnaryServerInterceptor{
		unaryInterceptor,
	}

	s := grpc.NewServer(grpc.ChainUnaryInterceptor(unary...))

	customer := server.New()

	pb.RegisterCustomersServiceServer(s, customer)

	println("Grpc port:4001")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func unaryInterceptor(ctx context.Context,
	req interface{},
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler) (interface{}, error) {
	start := time.Now()

	h, err := handler(ctx, req)

	log.Printf("request - method: %s\tduration: %s\terror: %v\n",
		info.FullMethod,
		time.Since(start),
		err,
	)

	return h, err
}

func getConnection(uri string) (*grpc.ClientConn, error) {
	k := keepalive.ClientParameters{}

	options := []grpc.DialOption{
		grpc.WithInsecure(),
		grpc.WithKeepaliveParams(k),
		grpc.WithDefaultCallOptions(
			grpc.WaitForReady(false),
		),
	}

	conn, err := grpc.Dial(
		uri,
		options...,
	)

	if err != nil {
		return nil, err
	}

	return conn, nil
}
