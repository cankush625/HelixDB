package cmd

import (
	"HelixDB/common"
	"HelixDB/db"
	"bytes"
	"fmt"
	"time"
)

var WrongNumberOfArgumentsError = fmt.Errorf("wrong number of arguments")

func Get(command common.Cmd) ([]byte, error) {
	if len(command.Args) != 1 {
		return common.RespError("wrong number of arguments"), WrongNumberOfArgumentsError
	}
	var buffer bytes.Buffer
	value, err := GetValueFromMemory(command.Args[0])
	if err != nil {
		buffer.WriteString("$" + "-1")
	} else {
		buffer.WriteString("+" + value)
	}
	buffer.WriteString(common.Terminator)
	return []byte(buffer.String()), nil
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
