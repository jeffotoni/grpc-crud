package main

import (
	"context"
	"fmt"
	"log"
//	"strconv"
	"time"

	pb "github.com/jeffotoni/grpc-crud/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var Addr = "localhost:4001"

func main() {
	// id := flag.String("id", "", "document")
	// flag.Parse()
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))

	conn, err := grpc.Dial(Addr, opts...)
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}
	defer conn.Close()
	client := pb.NewCustomersServiceClient(conn)

	ticker := time.NewTicker(time.Millisecond * time.Duration(50))
	done := make(chan bool)
	i := 0
	go func(i int) {
		for {
			select {
			case <-done:
				return
			case t := <-ticker.C:
				i++
				fmt.Println("ticker:", t)
				//istr := strconv.Itoa(i)
				getReq := &pb.GetRequest{Document: "43009665923"}
				customer, err := client.Get(context.Background(), getReq)
				if err != nil {
					log.Fatalf("%v.Get(_) = _, %v: ", client, err)
				}
				log.Println(customer)
			}
		}
	}(i)

	time.Sleep(time.Second * time.Duration(60))
	ticker.Stop()
	done <- true
	println("done")
}
