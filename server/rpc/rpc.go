package rpc

import (
	"fmt"
	"log"
	"time"

	"github.com/yuuzin217/grpc-cli-chat/chatService/pb"
	"github.com/yuuzin217/grpc-cli-chat/server/room"
)

type server struct {
	pb.UnimplementedChatServiceServer
	connInfo *ConnectInfo
}

func NewServer() *server {
	return &server{
		connInfo: NewConnectInfo(),
	}
}

// マッチング
func (s *server) Matching(req *pb.MatchingRequest, stream pb.ChatService_MatchingServer) error {

	// ユーザー情報登録
	u := room.NewUser(req.Name, stream)

	// 待機ルームに追加
	selfIdx := AddUser(u)

	// マッチング相手を探索（見つかるまで）
	for {
		Matching(selfIdx)
		if err := stream.Send(
			&pb.MatchingResponse{IsConnected: false, NoticeMessage: "Waiting to be matched..."},
		); err != nil {
			return err
		}
		time.Sleep(3 * time.Second)
	}
}

// チャット通信
func (s *server) Connect(stream pb.ChatService_ConnectServer) error {

	log.Println("chat connection was started")

	go s.connInfo.Receive(stream)
	go s.connInfo.SendMessage()

	for {
		select {
		case err := <-s.connInfo.errChan:
			// TODO: エラーによって適切な処理を行う
			fmt.Println(err)
		default:
			continue
		}
	}
}
