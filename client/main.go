package main

import (
	"bufio"
	"context"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/yuuzin217/grpc-cli-chat/chatService/pb"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
)

var addr = flag.String("addr", "localhost:50051", "the address to connect to")

type client struct {
	conn *grpc.ClientConn
	pb.ChatServiceClient
	sessionID string
	roomID    string
	name      string
}

func newClient() *client {
	conn, err := grpc.NewClient(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalln("did not connect:", err)
	}
	return &client{conn, pb.NewChatServiceClient(conn), "", "", ""}
}

func (c *client) setup(ctx context.Context) error {
	fmt.Println("welcome to grpc-cli-chat!!")

	// 名前入力
	var name string
	fmt.Print("enter your name... > ")
	fmt.Scan(&name)

	// マッチング
	stream, err := c.Matching(ctx, &pb.MatchingRequest{Name: name})
	if err != nil {
		return err
	}
	for {
		resp, err := stream.Recv()
		if err != nil {
			return err
		}
		if resp.IsConnected {
			c.name = name
			c.roomID = resp.RoomID
			c.sessionID = resp.SessionID
			fmt.Println(resp.NoticeMessage)
			return nil
		}
		fmt.Println(resp.NoticeMessage)
	}
}

func (c *client) connect(ctx context.Context) error {
	md := metadata.Pairs("session_id", c.sessionID)
	ctx = metadata.NewOutgoingContext(ctx, md)
	// チャット開始
	stream, err := c.Connect(ctx)
	if err != nil {
		return fmt.Errorf("failed to grpc 'Connect': %v", err)
	}
	defer stream.CloseSend()
	// コネクション情報をサーバー側に登録するため空リクエスト
	if err := stream.Send(&pb.SendMessage{RoomID: c.roomID}); err != nil {
		return err
	}
	errChan := make(chan error)
	go c.sendMsg(errChan, stream)
	go recvMsg(errChan, stream)
	for {
		select {
		case err := <-errChan:
			// TODO: エラーによって適切な処理を行う
			fmt.Println(err)
		default:
			continue
		}
	}
}

func (c *client) sendMsg(errCh chan<- error, stream pb.ChatService_ConnectClient) {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		scanner.Scan()
		msg := scanner.Text()
		if msg != "" {
			if err := stream.Send(&pb.SendMessage{
				RoomID:  c.roomID,
				Name:    c.name,
				Content: msg,
			}); err != nil {
				errCh <- err
			}
		}
	}
}

func recvMsg(errCh chan<- error, stream pb.ChatService_ConnectClient) {
	for {
		recv, err := stream.Recv()
		if err != nil {
			errCh <- err
		}
		fmt.Println(fmt.Sprint(recv.Name, ": ", recv.Content))
	}
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
