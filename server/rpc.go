package main

import (
	"context"
	"fmt"
	"log"

	"github.com/yuuzin217/grpc-cli-chat/chatService/pb"
	"github.com/yuuzin217/grpc-cli-chat/pkg/room"
	"google.golang.org/protobuf/types/known/emptypb"
)

// ルーム一覧取得
func (*server) GetRoomList(ctx context.Context, _ *emptypb.Empty) (*pb.RoomListResponse, error) {
	return &pb.RoomListResponse{RoomList: room.RoomList}, nil
}

// ルーム参加
func (s *server) JoinRoom(ctx context.Context, req *pb.JoinRequest) (*pb.JoinResponse, error) {
	// validate
	if _, isExists := room.RoomList[req.RoomNumber]; !isExists {
		return nil, fmt.Errorf("specified room number does not exist")
	}
	if req.Name == "" {
		return nil, fmt.Errorf("No name entered")
	}
	// ルーム参加
	r := room.GetOrCreateRoom(int(req.RoomNumber))
	u, err := room.NewUser(req.Name)
	if err != nil {
		return nil, err
	}
	r.AddUser(u)
	return &pb.JoinResponse{UserID: u.ID}, nil
}

// チャット通信
func (s *server) Connect(stream pb.ChatService_ConnectServer) error {
	log.Println("chat connection was started")
	msgChan := make(chan broadcastMsg)
	errChan := make(chan error)
	go s.recvMsg(errChan, msgChan, stream)
	go s.sendMsg(errChan, msgChan)
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

type broadcastMsg struct {
	roomNum    int
	senderId   string
	senderName string
	msg        string
}

func sendBroadcastMsg(msgChan chan<- broadcastMsg, roomNum int, id string, name string, msg string) error {
	msgChan <- broadcastMsg{
		roomNum:    roomNum,
		senderId:   id,
		senderName: name,
		msg:        msg,
	}
	return writeLog(roomNum, id, name, msg)
}

func (s *server) recvMsg(errChan chan<- error, msgChan chan<- broadcastMsg, stream pb.ChatService_ConnectServer) {
	for {
		req, err := stream.Recv()
		if err != nil {
			errChan <- err
		}
		mInfo := req.MessageInfo
		r, err := room.GetRoom(int(mInfo.RoomNumber))
		if err != nil {
			errChan <- err
		}
		u, ok := r.Users[mInfo.SenderID]
		if !ok {
			errChan <- fmt.Errorf("does not exists user: user id (%s), name (%s)", mInfo.SenderID, mInfo.SenderName)
		}

		msg := mInfo.Content
		if !u.IsSetCache {
			u.SetStreamCache(stream)
			msg = "[INFO] joined in room!!"
		}
		if err := sendBroadcastMsg(
			msgChan, int(mInfo.RoomNumber), mInfo.SenderID, mInfo.SenderName, msg,
		); err != nil {
			errChan <- err
		}
	}
}

func (s *server) sendMsg(errChan chan<- error, msgChan <-chan broadcastMsg) {
	for {
		select {
		case broadcast := <-msgChan:
			r, err := room.GetRoom(broadcast.roomNum)
			if err != nil {
				errChan <- err
				continue
			}
			for _, user := range r.Users {
				if user.ID == broadcast.senderId {
					// 送信者自身にはエコーバックしない
					continue
				}
				if user.StreamCache == nil {
					// （ないとは思うが）stream が存在しないので無視
					continue
				}
				if err := user.StreamCache.Send(
					&pb.ConnectResponse{
						MessageInfo: &pb.MessageInfo{
							SenderID:   broadcast.senderId,
							SenderName: broadcast.senderName,
							Content:    broadcast.msg,
						},
					},
				); err != nil {
					errChan <- err
				}
			}
		default:
			continue
		}
	}
}
