package main

import (
	"flag"
	"fmt"
	"log"
	"net"

	"github.com/yuuzin217/grpc-cli-chat/chatService/pb"
	"github.com/yuuzin217/grpc-cli-chat/server/rpc"
	"google.golang.org/grpc"
)

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
	s := rpc.NewServer()
	gs := grpc.NewServer()
	pb.RegisterChatServiceServer(gs, s)
	if err := gs.Serve(listen); err != nil {
		log.Fatalln("failed to serve:", err)
	}
}
