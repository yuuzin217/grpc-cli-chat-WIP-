package main

import (
	"context"
	"fmt"
	"io"
	"log"

	"github.com/yuuzin217/grpc-cli-chat/chatService/pb"
	"golang.org/x/sync/errgroup"

	"github.com/google/uuid"
)

var roomList map[int]string = map[int]string{
	1: "room1",
	2: "room2",
	3: "room3",
}

func (*server) GetRoomList(ctx context.Context, req *pb.RoomListRequest) (*pb.RoomListResponse, error) {
	res := &pb.RoomListResponse{}
	for i, room := range roomList {
		res.RoomList = append(res.RoomList, fmt.Sprint(i, ": ", room))
	}
	return res, nil
}

type broadcastMsg struct {
	roomNum int
	from    string
	msg     string
}

func (s *server) JoinRoom(ctx context.Context, req *pb.JoinRequest) (*pb.JoinResponse, error) {
	id, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}
	userId := UserID(id.String())
	s.clients[userId] = &client{
		name:        req.Name,
		joinRoomNum: int(req.RoomNumber),
	}
	return &pb.JoinResponse{
		UserID: id.String(),
	}, nil
}

func (s *server) Connect(stream pb.ChatService_ConnectServer) error {
	log.Println("chat connection was started")
	var eg errgroup.Group
	msgCh := make(chan broadcastMsg)
	defer close(msgCh)
	eg.Go(func() error {
		for {
			req, err := stream.Recv()
			if err == io.EOF {
				// to do.
				continue
			}
			if err != nil {
				return err
			}
			fromId := UserID(req.UserID)
			from, ok := s.clients[fromId]
			switch {
			case !ok:
				return fmt.Errorf("user is not created")
			case from == nil:
				return fmt.Errorf("not joined any room")
			default:
				// nothing to do.
			}
			if req.RegisterStream {
				s.clients[fromId].stream = stream
				continue
			}
			msgCh <- broadcastMsg{
				roomNum: from.joinRoomNum,
				from:    from.name,
				msg:     req.Message,
			}
		}
	})
	eg.Go(func() error {
		for {
			select {
			case broadcast := <-msgCh:
				for _, c := range s.clients {
					if c.joinRoomNum == broadcast.roomNum && c.name != broadcast.from {
						if err := c.stream.Send(
							&pb.ConnectResponse{Name: broadcast.from, Message: broadcast.msg},
						); err != nil {
							return err
						}
					}
				}
			default:
				continue
			}
		}
	})
	if err := eg.Wait(); err != nil {
		log.Println(err)
	}
	return fmt.Errorf("server stopped")
}
