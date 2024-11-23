package main

import (
	"flag"
	"fmt"
	"grpc-cli-chat/chatService/pb"
	"log"
	"net"

	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedChatServiceServer
}

var port = flag.Int("port", 50051, "the port to serve on")

func init() {
	flag.Parse()
}

func main() {
	listen, err := net.Listen("tcp", fmt.Sprintf("localhost:%v", *port))
	if err != nil {
		log.Fatalln("failed to listen:", err)
	}
	fmt.Println("server listening at", listen.Addr())

	s := grpc.NewServer()
	pb.RegisterChatServiceServer(s, &server{})
	if err := s.Serve(listen); err != nil {
		log.Fatalln("failed to serve:", err)
	}
}
