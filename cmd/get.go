package cmd

import (
	"HelixDB/common"
	"HelixDB/db"
	"fmt"
	"time"
)

var WrongNumberOfArgumentsError = fmt.Errorf("wrong number of arguments")

func Get(command common.Cmd) ([]byte, error) {
	if len(command.Args) != 1 {
		return common.RespError("wrong number of arguments"), WrongNumberOfArgumentsError
	}
	value, err := GetValueFromMemory(command.Args[0])
	if err != nil {
		return common.RespNullBulkString(), nil
	}
	return common.RespBulkString(value), nil
}

func GetValueFromMemory(key string) (string, error) {
	// Check if the key is expired
	expirationTime, ok := db.KeyTTL.Load(key)
	if !ok {
		return "", fmt.Errorf("key not found")
	}
	currentTime := time.Now().UnixMilli()
	if expirationTime != nil && expirationTime.(int64) < currentTime {
		return "", fmt.Errorf("key expired")
	}

	data, ok := db.DB.Load(key)
	if !ok {
		return "", fmt.Errorf("key not found")
	}
	return data.(string), nil
}
