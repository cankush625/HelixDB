package cmd

import (
	"HelixDB/common"
	"HelixDB/db"
	"bytes"
	"fmt"
)

var WrongNumberOfArgumentsError = fmt.Errorf("wrong number of arguments")

func Get(command []string) ([]byte, error) {
	if len(command) != 2 {
		return common.RespError("wrong number of arguments"), WrongNumberOfArgumentsError
	}
	var buffer bytes.Buffer
	value, err := GetValueFromMemory(command[1])
	if err != nil {
		buffer.WriteString("$" + "-1")
	} else {
		buffer.WriteString("+" + value)
	}
	buffer.WriteString(common.Terminator)
	return []byte(buffer.String()), nil
}

func GetValueFromMemory(key string) (string, error) {
	data, ok := db.DB.Load(key)
	if !ok {
		return "", fmt.Errorf("key not found")
	}
	return data.(string), nil
}
