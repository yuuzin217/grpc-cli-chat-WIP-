package main

import (
	"flag"
	"fmt"
	"log"
	"net"

	"github.com/yuuzin217/grpc-cli-chat/chatService/pb"
	"google.golang.org/grpc"
)

type UserID string // UUIDv4

type client struct {
	name        string
	joinRoomNum int
	stream      pb.ChatService_SendAndUpdateServer
}

type server struct {
	pb.UnimplementedChatServiceServer
	clients map[UserID]*client
}

var port = flag.Int("port", 50051, "the port to serve on")

func init() {
	flag.Parse()
}

func newServer() *server {
	return &server{
		clients: make(map[UserID]*client),
	}
}

func main() {
	listen, err := net.Listen("tcp", fmt.Sprintf("localhost:%v", *port))
	if err != nil {
		log.Fatalln("failed to listen:", err)
	}
	fmt.Println("server listening at", listen.Addr())
	s := grpc.NewServer()
	pb.RegisterChatServiceServer(s, newServer())
	if err := s.Serve(listen); err != nil {
		log.Fatalln("failed to serve:", err)
	}
}
