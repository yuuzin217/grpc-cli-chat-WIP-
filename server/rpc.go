package main

import (
	"context"
	"fmt"
	"io"
	"log"

	"github.com/yuuzin217/grpc-cli-chat/chatService/pb"

	"github.com/google/uuid"
)

var roomList map[int32]string = map[int32]string{
	1: "room1",
	2: "room2",
	3: "room3",
}

// ルーム一覧取得
func (*server) GetRoomList(ctx context.Context, req *pb.RoomListRequest) (*pb.RoomListResponse, error) {
	return &pb.RoomListResponse{RoomList: roomList}, nil
}

// ルーム参加
func (s *server) JoinRoom(ctx context.Context, req *pb.JoinRequest) (*pb.JoinResponse, error) {
	// validate
	if _, isExists := roomList[req.RoomNumber]; !isExists {
		return nil, fmt.Errorf("specified room number does not exist")
	}
	if req.Name == "" {
		return nil, fmt.Errorf("No name entered")
	}
	// generate id
	id, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}
	userId := id.String()
	s.clients[userId] = &client{
		name:        req.Name,
		joinRoomNum: int(req.RoomNumber),
	}
	return &pb.JoinResponse{UserID: userId}, nil
}

// チャット通信
func (s *server) Connect(stream pb.ChatService_ConnectServer) error {
	log.Println("chat connection was started")
	msgChan := make(chan broadcastMsg)
	errChan := make(chan error)
	defer close(msgChan)
	defer close(errChan)
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
	roomNum int
	fromId  string
	from    string
	msg     string
}

func sendBroadcastMsg(msgChan chan<- broadcastMsg, roomNum int, id string, name string, msg string) error {
	msgChan <- broadcastMsg{
		roomNum: roomNum,
		fromId:  id,
		from:    name,
		msg:     msg,
	}
	return writeLog(roomNum, id, name, msg)
}

func (s *server) recvMsg(errChan chan<- error, msgChan chan<- broadcastMsg, stream pb.ChatService_ConnectServer) {
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			continue
		}
		if err != nil {
			errChan <- err
		}
		client, err := s.getClient(req.UserID)
		if err != nil {
			errChan <- err
		}
		msg := req.Message
		if req.RegisterConnection {
			s.clients[req.UserID].stream = stream
			msg = "[INFO] joined in room!!"
		}
		if err := sendBroadcastMsg(
			msgChan, client.joinRoomNum, req.UserID, client.name, msg,
		); err != nil {
			errChan <- err
		}
	}
}

func (s *server) getClient(userId string) (*client, error) {
	commonErrStr := fmt.Sprintf("getClient error: user id (%s)", userId)
	client, ok := s.clients[userId]
	switch {
	case !ok:
		return nil, fmt.Errorf("client is not registered: %s", commonErrStr)
	case client == nil:
		return nil, fmt.Errorf("not joined any room: %s", commonErrStr)
	default:
		return client, nil
	}
}

func (s *server) sendMsg(errChan chan<- error, msgChan <-chan broadcastMsg) {
	for {
		select {
		case broadcast := <-msgChan:
			for _, target := range s.getSendTargets(broadcast.roomNum, broadcast.fromId) {
				if err := target.stream.Send(
					&pb.ConnectResponse{Name: broadcast.from, Message: broadcast.msg},
				); err != nil {
					errChan <- err
				}
			}
		default:
			continue
		}
	}
}

func (s *server) getSendTargets(targetRoomNum int, fromId string) map[string]*client {
	copied := make(map[string]*client)
	for k, v := range s.clients {
		if targetRoomNum == v.joinRoomNum {
			copied[k] = v
		}
	}
	// 送信元へはエコーされなくていいので除外
	// ※ map を新たに作り直しているので実質ディープコピーになっており元データには影響ない
	delete(copied, fromId)
	return copied
}
