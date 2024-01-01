package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"

	pbx "github.com/gerry-sabar/go-hello/proto"
)

var port = 50051

type server struct {
	pbx.UnimplementedChatServer
}

func (s *server) SayHello(ctx context.Context, in *pbx.ChatRequest) (*pbx.ChatReply, error) {
	log.Printf("Received: %v", in.GetMessage())
	return &pbx.ChatReply{Message: "Hello " + in.GetMessage()}, nil
}

func main() {
	fmt.Printf("server is running... \n")
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))

	if err != nil {
		log.Fatalf("failed to listen")
	}

	s := grpc.NewServer()
	pbx.RegisterChatServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve")
	}
}
