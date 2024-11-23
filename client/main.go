package main

import (
	"context"
	"flag"
	"fmt"
	"log"

	"github.com/yuuzin217/grpc-cli-chat/chatService/pb"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var addr = flag.String("addr", "localhost:50051", "the address to connect to")

type client struct {
	conn *grpc.ClientConn
	pb.ChatServiceClient
}

func newClient() *client {
	conn, err := grpc.NewClient(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalln("did not connect:", err)
	}
	return &client{conn, pb.NewChatServiceClient(conn)}
}

func (c *client) setup(ctx context.Context) error {
	resp, err := c.JoinRoom(ctx, &pb.JoinRequest{Name: "aaa", RoomNumber: 2})
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(resp)
	fmt.Println("welcome to grpc-cli-chat")
	fmt.Println("pls choose chat room:")
	res, err := c.GetRoomList(ctx, nil)
	if err != nil {
		return err
	}
	for _, room := range res.RoomList {
		fmt.Println(room)
	}
	var roomNum int
	fmt.Print(" > ")
	fmt.Scan(&roomNum)
	fmt.Printf("your choice room number: %v", roomNum)
	return nil
}

func main() {
	flag.Parse()
	c := newClient()
	defer c.conn.Close()

	ctx := context.Background()

	if err := c.setup(ctx); err != nil {
		log.Fatalln("failed to setup:", err)
	}
}
