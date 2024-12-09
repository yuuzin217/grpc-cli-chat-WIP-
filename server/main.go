package main

import (
	"flag"
	"fmt"
	"log"
	"net"

	"github.com/yuuzin217/grpc-cli-chat/chatService/pb"
	"google.golang.org/grpc"
)

var port = flag.Int("port", 50051, "the port to serve on")

func init() {
	flag.Parse()
}

type client struct {
	name        string
	joinRoomNum int
	stream      pb.ChatService_ConnectServer
}

type server struct {
	pb.UnimplementedChatServiceServer
	clients map[string]*client // map[UserID]*client
}

func newServer() *server {
	return &server{
		clients: make(map[string]*client),
	}
}

func main() {
	listen, err := net.Listen("tcp", fmt.Sprintf(":%v", *port))
	if err != nil {
		log.Fatalln("failed to listen:", err)
	}
	fmt.Println("server listening at", listen.Addr())
	s := newServer()
	gs := grpc.NewServer()
	pb.RegisterChatServiceServer(gs, s)
	if err := gs.Serve(listen); err != nil {
		log.Fatalln("failed to serve:", err)
	}
}

// // https://pkg.go.dev/google.golang.org/grpc#StreamServerInterceptor
// func streamInterceptor(srv any, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
// 	s, ok := srv.(*server)
// 	if !ok {
// 		return fmt.Errorf("structure mismatch: expected %T / actually %T", &server{}, srv)
// 	}

// 	err := handler(srv, ss)

// 	// NOTE: 後処理が必要であればここに追記する

// 	return err
// }

// // func preProcess() {}
