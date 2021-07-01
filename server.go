// +build server

package main

import (
	"context"
	"demo/demo/pb"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

type Object struct {
	pb.SendServer
}

func (this *Object) Do(ctx context.Context, msg *wrapperspb.StringValue) (*emptypb.Empty, error) {
	log.Println("it call")
	fmt.Printf("recived:%v\n", msg.Value)
	return &emptypb.Empty{}, nil
}

func main() {
	tcp, err := net.Listen("tcp", "0.0.0.0:9999")
	if err != nil {
		log.Fatal(err)
	}

	server := grpc.NewServer()
	pb.RegisterSendServer(server, &Object{})
	log.Println("servering")
	server.Serve(tcp)
}
