package chatlog

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/yuuzin217/grpc-cli-chat/server/room"
)

const LOG_COMMON_PATH = "server/log/"

// ディレクトリが存在するかチェック、なければ作成する
func isExistsDir(path string) error {
	if err := os.MkdirAll(path, 0755); err != nil {
		return fmt.Errorf("mkdir unknown error: %v", err)
	}
	return nil
}

func getFileName(roomID room.RoomID) (string, error) {
	date := time.Now().Format(time.DateOnly)
	path := LOG_COMMON_PATH + date
	if err := isExistsDir(path); err != nil {
		return "", err
	}
	return fmt.Sprintf("%s/roomID_%s.log", path, roomID), nil
}

func WriteLog(roomID room.RoomID, name string, msg string) error {
	var logFile *os.File
	defer logFile.Close()
	fileName, err := getFileName(roomID)
	if err != nil {
		return err
	}
	_, err = os.Stat(fileName)
	switch {
	case os.IsNotExist(err):
		// 新規作成
		if logFile, err = os.Create(fileName); err != nil {
			return err
		}
	case !os.IsNotExist(err):
		// 追記モードで開く
		if logFile, err = os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY, 0644); err != nil {
			return err
		}
	default:
		// それ以外のエラー
		if err != nil {
			return fmt.Errorf("file exists check unknown error: %v", err)
		}
	}
	log.SetOutput(logFile)
	// 標準出力にも残したい場合はこちら
	// log.SetOutput(io.MultiWriter(logFile, os.Stdout))
	log.Printf(": Name: %s, Message: %s\n", name, msg)
	log.SetOutput(os.Stdout)
	return nil
}
