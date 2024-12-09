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

type broadcastMsg struct {
	roomNum int
	from    string
	fromId  string
	msg     string
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
			fmt.Println(err)
		default:
			continue
		}
	}
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
		from, ok := s.clients[req.UserID]
		switch {
		case !ok:
			errChan <- fmt.Errorf("user is not created")
		case from == nil:
			errChan <- fmt.Errorf("not joined any room")
		default:
			// nothing to do.
		}
		id := req.UserID
		num := from.joinRoomNum
		name := from.name
		msg := req.Message
		if req.RegisterStream {
			s.clients[id].stream = stream
			continue
		}
		if err := writeLog(num, id, name, msg); err != nil {
			errChan <- err
		}
		msgChan <- broadcastMsg{
			roomNum: num,
			from:    name,
			fromId:  id,
			msg:     msg,
		}
	}
}

func (s *server) sendMsg(errChan chan<- error, msgChan <-chan broadcastMsg) {
	for {
		select {
		case broadcast := <-msgChan:
			for id, client := range s.clients {
				if client.joinRoomNum == broadcast.roomNum && id != broadcast.fromId {
					if err := client.stream.Send(
						&pb.ConnectResponse{Name: broadcast.from, Message: broadcast.msg},
					); err != nil {
						errChan <- err
					}
				}
			}
		default:
			continue
		}
	}
}
