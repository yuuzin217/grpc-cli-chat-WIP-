package room

import (
	"fmt"
	"sync"
)

var RoomList map[int32]string = map[int32]string{
	1: "room1",
	2: "room2",
	3: "room3",
}

var rooms = make(map[int]*Room)

func GetRooms() map[int]*Room {
	return rooms
}

type broadcastMsg struct {
	senderId string
	msg      string
}

type Room struct {
	sync.RWMutex
	msgCh chan *broadcastMsg
	//    map[UserID]*User
	Users map[string]*User
}

func newRoom() *Room {
	return &Room{
		Users: make(map[string]*User),
	}
}

// ルーム情報を取得、なければ作成
func GetOrCreateRoom(num int) *Room {
	room, isExists := rooms[num]
	if !isExists {
		rooms[num] = newRoom()
		return rooms[num]
	}
	return room
}

// ルーム情報を取得、なければエラー
func GetRoom(num int) (*Room, error) {
	room, isExists := rooms[num]
	if !isExists {
		return nil, fmt.Errorf("Room does not exist: number (%d)", num)
	}
	return room, nil
}

func (r *Room) AddUser(u *User) {
	r.RWMutex.Lock()
	{
		r.Users[u.ID] = u
	}
	r.RWMutex.Unlock()
}

func (r *Room) ExitUser(id string) {
	r.RWMutex.Lock()
	{
		delete(r.Users, id)
	}
	r.RWMutex.Unlock()
}

func (r *Room) GetUsers() map[string]*User {
	var users map[string]*User
	r.RWMutex.RLock()
	{
		users = r.Users
	}
	r.RWMutex.Unlock()
	return users
}
