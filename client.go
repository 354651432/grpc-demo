package main

import (
	"context"
	"log"

	"demo/demo/pb"

	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

func main() {
	conn, err := grpc.Dial("127.0.0.1:9999", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	client := pb.NewSendClient(conn)
	client.Do(context.Background(), &wrapperspb.StringValue{Value: "message from a dual client"})
}
