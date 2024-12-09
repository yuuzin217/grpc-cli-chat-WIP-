package main

import (
	"fmt"
	"log"
	"os"
	"time"
)

const LOG_COMMON_PATH = "server/log/"

func mkdir(path string) error {
	_, err := os.Stat(path)
	switch {
	case os.IsNotExist(err):
		return os.Mkdir(path, 0644)
	case !os.IsNotExist(err):
		return nil
	default:
		if err != nil {
			return fmt.Errorf("mkdir unknown error: %v", err)
		}
	}
	return nil
}

func getFileName(roomNum int) (string, error) {
	t := time.Now()
	var logPath string
	switch roomNum {
	case 1:
		logPath = LOG_COMMON_PATH + "room1"
		if err := mkdir(logPath); err != nil {
			return "", err
		}
	case 2:
		logPath = LOG_COMMON_PATH + "room2"
		if err := mkdir(logPath); err != nil {
			return "", err
		}
	case 3:
		logPath = LOG_COMMON_PATH + "room3"
		if err := mkdir(logPath); err != nil {
			return "", err
		}
	default:
		logPath = LOG_COMMON_PATH + "unknown"
		if err := mkdir(logPath); err != nil {
			return "", err
		}
	}
	return fmt.Sprintf("%s/%s.log", logPath, t.Format(time.DateOnly)), nil
}

func writeLog(roomNum int, userId string, name string, msg string) error {
	fileName, err := getFileName(roomNum)
	if err != nil {
		return err
	}
	var logFile *os.File
	defer logFile.Close()
	_, err = os.Stat(fileName)
	switch {
	case os.IsNotExist(err):
		if logFile, err = os.Create(fileName); err != nil {
			return err
		}
	case !os.IsNotExist(err):
		if logFile, err = os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY, 0644); err != nil {
			return err
		}
	default:
		if err != nil {
			return fmt.Errorf("file exists check unknown error: %v", err)
		}
	}
	log.SetOutput(logFile)
	log.Println(fmt.Sprintf(": UserID: %s, Name: %s, Message: %s", userId, name, msg))
	log.SetOutput(os.Stdout)
	return nil
}
