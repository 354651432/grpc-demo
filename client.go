package main

import (
	"context"
	"log"
	"time"

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
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	client.Do(ctx, &wrapperspb.StringValue{Value: "来点中文 ok 吗"})
}
