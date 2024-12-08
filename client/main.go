package main

import (
	"bufio"
	"context"
	"flag"
	"fmt"
	"log"
	"maps"
	"os"
	"slices"

	"github.com/yuuzin217/grpc-cli-chat/chatService/pb"
	"golang.org/x/sync/errgroup"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var addr = flag.String("addr", "localhost:50051", "the address to connect to")

type client struct {
	conn *grpc.ClientConn
	pb.ChatServiceClient
	stream pb.ChatService_ConnectClient
	userId string
	name   string
}

func newClient() *client {
	conn, err := grpc.NewClient(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalln("did not connect:", err)
	}
	return &client{conn, pb.NewChatServiceClient(conn), nil, "", ""}
}

func (c *client) setup(ctx context.Context) error {
	fmt.Println("welcome to grpc-cli-chat!!")
	var roomNum int
	var name string
	res, err := c.GetRoomList(ctx, nil)
	if err != nil {
		return err
	}
	// ルーム選択
	for {
		fmt.Println("select a chat room number:")
		for _, key := range slices.Sorted(maps.Keys(res.RoomList)) {
			fmt.Println(fmt.Sprint(key, ": ", res.RoomList[key]))
		}
		fmt.Print(" > ")
		fmt.Scan(&roomNum)
		if _, b := res.RoomList[int32(roomNum)]; !b {
			fmt.Println("Unknown room number. Please select again")
			continue
		} else {
			break
		}
	}
	// 名前入力
	for {
		fmt.Printf("Your selected chat room number is %v\n", roomNum)
		fmt.Print("enter your name... > ")
		fmt.Scan(&name)
		if name == "" {
			fmt.Println("name has not been entered. Please re-enter")
			continue
		} else {
			break
		}
	}
	// ルーム参加
	{
		res, err := c.JoinRoom(ctx, &pb.JoinRequest{Name: name, RoomNumber: int32(roomNum)})
		if err != nil {
			return err
		}
		c.name = name
		c.userId = res.UserID
	}
	fmt.Println("welcome!!", name)
	fmt.Println("------ Let's talking !! ------", "\n")
	return nil
}

func (c *client) connect(ctx context.Context) error {
	// チャット開始
	stream, err := c.Connect(ctx)
	if err != nil {
		return fmt.Errorf("failed to grpc 'Connect': %v", err)
	}
	defer stream.CloseSend()
	// stream をサーバー側に登録するため空リクエスト
	if err := stream.Send(&pb.ConnectRequest{
		UserID:         c.userId,
		RegisterStream: true,
	}); err != nil {
		return err
	}
	var eg errgroup.Group
	scanner := bufio.NewScanner(os.Stdin)
	eg.Go(func() error {
		for {
			scanner.Scan()
			msg := scanner.Text()
			if msg != "" {
				if err := stream.Send(&pb.ConnectRequest{
					UserID:  c.userId,
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
		// TODO: ログレベル等によって挙動を切り分ける
		return err
	}
	return fmt.Errorf("unknown Error")
}

func main() {
	flag.Parse()
	c := newClient()
	defer c.conn.Close()
	ctx := context.Background()
	if err := c.setup(ctx); err != nil {
		log.Fatalln("failed to setup:", err)
	}
	if err := c.connect(ctx); err != nil {
		log.Fatalln("failed to connect:", err)
	}
}
