package rpc

import (
	"github.com/yuuzin217/grpc-cli-chat/chatService/pb"
	chatlog "github.com/yuuzin217/grpc-cli-chat/server/chatLog"
	"github.com/yuuzin217/grpc-cli-chat/server/room"
	"google.golang.org/grpc/metadata"
)

func extractSessionID(stream pb.ChatService_ConnectServer) room.SessionID {
	// メタデータからセッションIDを取得
	md, ok := metadata.FromIncomingContext(stream.Context())
	if !ok {
		_ = ok
		// TODO: エラー返却
	}
	val, ok := md["session_id"]
	if !ok {
		_ = ok
		// TODO: エラー返却
	}
	return room.SessionID(val[0])
}

type ConnectInfo struct {
	roomInfo *room.ChatRoom
	msgChan  chan broadcastMessage
	errChan  chan error
}

func NewConnectInfo() *ConnectInfo {
	return &ConnectInfo{
		msgChan: make(chan broadcastMessage),
		errChan: make(chan error),
	}
}

type broadcastMessage struct {
	senderName  string
	content     string
	targetUsers []*room.User
}

func (c *ConnectInfo) Receive(stream pb.ChatService_ConnectServer) {
	sessionID := extractSessionID(stream)
	for {
		recv, err := stream.Recv()
		if err != nil {
			c.errChan <- err
		}
		// ルーム情報のキャッシュがあるか検索
		if c.roomInfo == nil {
			c.roomInfo = room.GetRoom(room.RoomID(recv.RoomID))
		}
		sender, err := c.roomInfo.GetUser(sessionID)
		if err != nil {
			// TODO: エラー処理
			c.errChan <- err
		}
		if !sender.IsSetConnectStream() {
			// ストリームキャッシュがない場合はセット
			sender.SetConnectStream(stream)
		}
		if recv.Content == "" {
			// メッセージがないので無視
			continue
		}
		targetUsers := []*room.User{}
		for id, u := range c.roomInfo.GetUsers() {
			if id == sessionID {
				// 送信者は無視
				continue
			}
			targetUsers = append(targetUsers, u)
		}
		c.msgChan <- broadcastMessage{
			senderName:  sender.GetName(),
			content:     recv.Content,
			targetUsers: targetUsers,
		}
	}
}

func (c *ConnectInfo) SendMessage() {
	for {
		select {
		case broadcastMsg := <-c.msgChan:
			for _, target := range broadcastMsg.targetUsers {
				if !target.IsSetConnectStream() {
					// stream がセットされていないので無視
					continue
				}
				if err := target.Send(
					broadcastMsg.senderName,
					broadcastMsg.content,
				); err != nil {
					c.errChan <- err
				}
			}
			if err := chatlog.WriteLog(
				c.roomInfo.GetRoomID(),
				broadcastMsg.senderName,
				broadcastMsg.content,
			); err != nil {
				c.errChan <- err
			}
		default:
			continue
		}
	}
}
