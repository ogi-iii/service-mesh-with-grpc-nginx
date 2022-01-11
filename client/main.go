package main

import (
	"context"
	"log"
	"service-mesh-with-grpc-nginx/protobuf"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

func main() {
	log.Println("Client started.")
	certFile := "../ssl/ca.crt"
	creds, sslErr := credentials.NewClientTLSFromFile(certFile, "")
	if sslErr != nil {
		log.Fatalln("failed to load certificates: ", sslErr)
		return
	}
	cc, err := grpc.Dial("127.0.0.1:50055", grpc.WithTransportCredentials(creds)) // using ssl
	// cc, err := grpc.Dial("localhost:50055", grpc.WithInsecure())               // dial to the server without secure ssl
	if err != nil {
		log.Fatalln(err)
	}
	defer cc.Close()

	c := protobuf.NewGreeterClient(cc) // create client

	ctx := context.Background()
	req := &protobuf.HelloRequest{
		Name: "Golang",
	}
	res, err := c.SayHello(ctx, req) // call Unary API
	if err != nil {
		log.Fatalln(err)
	}
	log.Printf("Unary: %v\n", res.GetReply())
}
