package main

import (
	"context"
	"fmt"
	"grpc-cli-chat/chatService/pb"
	"io"
	"log"

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

type connection struct {
	name    string
	roomNum int
}

var connections map[string]*connection = make(map[string]*connection)

func (*server) JoinRoom(ctx context.Context, req *pb.JoinRequest) (*pb.JoinResponse, error) {
	id, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}
	connections[id.String()] = &connection{
		name:    req.Name,
		roomNum: int(req.RoomNumber),
	}
	return &pb.JoinResponse{
		Name:       req.Name,
		RoomNumber: req.RoomNumber,
	}, nil
}

func (*server) SendAndUpdate(stream pb.ChatService_SendAndUpdateServer) error {
	log.Println("send and update was invoked")
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			// to do.
			continue
		}
		if err != nil {
			return err
		}
		stream.Send(
			&pb.ChatResponse{
				Name:    req.Name,
				Message: req.Message,
			},
		)
	}
}
