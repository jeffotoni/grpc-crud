package main

import (
	"context"
	"flag"
	"log"
	"time"

	pb "github.com/jeffotoni/grpc-crud/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var Addr = "localhost:4001"

func main() {
	id := flag.String("id", "", "document")
	flag.Parse()
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))

	conn, err := grpc.Dial(Addr, opts...)
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}
	defer conn.Close()
	client := pb.NewCustomersServiceClient(conn)

	log.Println("Get:")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	println("here:", *id)
	getReq := &pb.GetRequest{Document: *id}
	customer, err := client.Get(ctx, getReq)
	if err != nil {
		log.Fatalf("%v.Get(_) = _, %v: ", client, err)
	}
	log.Println(customer)
}
