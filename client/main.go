package main

import (
	"bufio"
	"context"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/yuuzin217/grpc-cli-chat/chatService/pb"
	"golang.org/x/sync/errgroup"

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

	// ルーム選択
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
	fmt.Printf("your choice room number: %v\n", roomNum)

	// 名前
	fmt.Print("pls your name > ")
	var name string
	fmt.Scan(&name)
	resp, err := c.JoinRoom(ctx, &pb.JoinRequest{Name: name, RoomNumber: int32(roomNum)})
	if err != nil {
		return err
	}
	fmt.Println("welcome!!", name)
	fmt.Println("lets talk!!")
	userId := resp.UserID

	// チャット開始
	stream, err := c.Connect(ctx)
	if err != nil {
		log.Fatalln(err)
	}
	defer stream.CloseSend()
	if err := stream.Send(&pb.ConnectRequest{
		UserID:         userId,
		RegisterStream: true,
	}); err != nil {
		log.Fatalln(err)
	}
	var eg errgroup.Group
	scanner := bufio.NewScanner(os.Stdin)
	eg.Go(func() error {
		for {
			scanner.Scan()
			msg := scanner.Text()
			if msg != "" {
				if err := stream.Send(&pb.ConnectRequest{
					UserID:  userId,
					Message: msg,
				}); err != nil {
					return err
				}
			}
		}
	})
	eg.Go(func() error {
		for {
			resp, err := stream.Recv()
			if err != nil {
				return err
			}
			fmt.Println(fmt.Sprint(resp.Name, ": ", resp.Message))
		}
	})
	if err := eg.Wait(); err != nil {
		log.Println(err)
	}
	return fmt.Errorf("chat end. bye")
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
