package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"service-mesh-with-grpc-nginx/protobuf"

	"google.golang.org/grpc"
)

type api struct {
	protobuf.UnimplementedGreeterServer // must be embedded
}

// define the process of Unary API
func (*api) SayHello(ctx context.Context, req *protobuf.HelloRequest) (*protobuf.HelloResponse, error) {
	log.Println("Unary called!")
	name := req.GetName()
	res := &protobuf.HelloResponse{
		Reply: fmt.Sprintf("Hello, %v!", name),
	}
	return res, nil
}

func main() {
	log.Println("api started.")

	lis, err := net.Listen("tcp", "0.0.0.0:50051") // create lister (default port: 50051)
	if err != nil {
		log.Fatalln(err)
	}
	s := grpc.NewServer()
	protobuf.RegisterGreeterServer(s, &api{}) // register the struct pointer as new-server

	if err = s.Serve(lis); err != nil { // serve the lister
		log.Fatalln(err)
	}
}
