package room

import (
	"github.com/yuuzin217/grpc-cli-chat/chatService/pb"
	"github.com/yuuzin217/grpc-cli-chat/server/common"
)

type SessionID string

type User struct {
	sessionID SessionID
	name      string
	mStream   *matchingStream
	cStream   *connectStream
}

type matchingStream struct {
	stream      pb.ChatService_MatchingServer
	isSetStream bool
}
type connectStream struct {
	stream      pb.ChatService_ConnectServer
	isSetStream bool
}

func NewUser(name string, stream pb.ChatService_MatchingServer) *User {
	return &User{
		sessionID: NewSessionID(),
		name:      name,
		mStream: &matchingStream{
			stream:      stream,
			isSetStream: true,
		},
		cStream: &connectStream{},
	}
}

func NewSessionID() SessionID {
	return SessionID(common.NewID())
}

func (u *User) GetSessionID() SessionID {
	return u.sessionID
}

func (u *User) GetName() string {
	return u.name
}

func (u *User) IsSetMatchingStream() bool {
	return u.mStream.isSetStream
}

func (u *User) IsSetConnectStream() bool {
	return u.cStream.isSetStream
}

func (u *User) SetMatchingStream(stream pb.ChatService_MatchingServer) {
	if u.mStream.isSetStream {
		return
	}
	u.mStream.stream = stream
	u.mStream.isSetStream = true
}

func (u *User) SetConnectStream(stream pb.ChatService_ConnectServer) {
	if u.cStream.isSetStream {
		return
	}
	u.cStream.stream = stream
	u.cStream.isSetStream = true
}

func (u *User) MatchingSend(isConnected bool, msg string, sessionID SessionID, roomID RoomID) error {
	if !u.IsSetMatchingStream() {
		// ユーザー作成時に stream も同時にセットするからここにくることはないはず...
		// TODO: エラー処理
		return nil
	}
	return u.mStream.stream.Send(
		&pb.MatchingResponse{
			IsConnected:   isConnected,
			NoticeMessage: msg,
			SessionID:     string(sessionID),
			RoomID:        string(roomID),
		},
	)
}

func (u *User) Send(senderName string, content string) error {
	if err := u.cStream.stream.Send(
		&pb.ReceivedMessage{
			Name:    u.name,
			Content: content,
		},
	); err != nil {
		return err
	}
	return nil
}
