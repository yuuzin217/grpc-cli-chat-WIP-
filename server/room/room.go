package room

import (
	"errors"
	"fmt"

	"github.com/yuuzin217/grpc-cli-chat/server/common"
)

type RoomID string

var ErrDoesNotExistsUser = errors.New("does not exists user")

type ChatRoom struct {
	roomID RoomID
	users  map[SessionID]*User
}

var ChatRooms = map[RoomID]*ChatRoom{}

func NewRoomID() RoomID {
	return RoomID(common.NewID())
}

func NewRoom(users ...*User) (RoomID, *ChatRoom) {
	for {
		roomID := NewRoomID()
		_, isExists := ChatRooms[roomID]
		if isExists {
			// ルーム ID 重複なので採番し直し
			continue
		} else {
			ChatRooms[roomID] = &ChatRoom{
				roomID: roomID,
				users:  make(map[SessionID]*User),
			}
			for _, user := range users {
				ChatRooms[roomID].users[user.GetSessionID()] = user
			}
			return roomID, ChatRooms[roomID]
		}
	}
}

func GetRoom(roomID RoomID) *ChatRoom {
	room := ChatRooms[roomID]
	// TODO: ルームがなかった場合の処理
	return room
}

func (r *ChatRoom) GetRoomID() RoomID {
	return r.roomID
}

func (r *ChatRoom) GetUsers() map[SessionID]*User {
	return r.users
}

func (r *ChatRoom) GetUser(id SessionID) (*User, error) {
	u, ok := r.users[id]
	if !ok {
		return nil, fmt.Errorf("%s: %s", ErrDoesNotExistsUser, fmt.Errorf("sessionID (%s)", id))
	}
	return u, nil
}
