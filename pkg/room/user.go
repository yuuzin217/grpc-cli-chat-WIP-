package room

import (
	"github.com/google/uuid"
	"github.com/yuuzin217/grpc-cli-chat/chatService/pb"
)

type User struct {
	ID          string
	Name        string
	StreamCache pb.ChatService_ConnectServer
	IsSetCache  bool
}

func NewUserID() (string, error) {
	// generate UUID v4
	id, err := uuid.NewRandom()
	if err != nil {
		return "", err
	}
	return id.String(), nil
}

func NewUser(name string) (*User, error) {
	userID, err := NewUserID()
	if err != nil {
		return nil, err
	}
	return &User{ID: userID, Name: name}, nil
}

func (u *User) SetStreamCache(stream pb.ChatService_ConnectServer) {
	u.StreamCache = stream
	u.IsSetCache = true
}
