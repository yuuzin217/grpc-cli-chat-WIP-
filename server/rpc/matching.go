package rpc

import (
	"math/rand/v2"
	"sort"
	"sync"

	"github.com/yuuzin217/grpc-cli-chat/server/room"
)

type waitingRoom struct {
	mu    sync.RWMutex
	users []*room.User
}

var wr *waitingRoom = &waitingRoom{}

func GetWaitingRoom() *waitingRoom {
	return wr
}

func AddUser(u *room.User) int {
	wr.mu.Lock()
	defer wr.mu.Unlock()
	{
		wr.users = append(wr.users, u)
		return len(wr.users) - 1
	}
}

func Matching(selfIdx int) {
	wr.mu.Lock()
	defer wr.mu.Unlock()
	{
		if len(wr.users) < 1 {
			// 待機に誰もいないので何もしない
			return
		}
		targetIdx := rand.IntN(len(wr.users))
		if targetIdx == selfIdx {
			// 自分自身なので無視
			return
		}
		self := wr.users[selfIdx]
		target := wr.users[targetIdx]

		// ルーム作成
		roomID, _ := room.NewRoom(self, target)

		// マッチング成功を通知
		for _, u := range wr.users {
			if err := u.MatchingSend(true, "Your match has been found!!", u.GetSessionID(), roomID); err != nil {
				// TODO: エラー処理
				_ = err
			}
		}

		// 要素を削除
		removeIndices := []int{targetIdx, selfIdx}
		sort.Ints(removeIndices)
		for i, idx := range removeIndices {
			if 0 < i {
				idx--
			}
			wr.users = append(wr.users[:idx], wr.users[idx+1:]...)
		}
	}
}
