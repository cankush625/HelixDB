package cmd

import (
	"HelixDB/common"
	"HelixDB/db"
	"bytes"
	"fmt"
)

func Get(command []string) ([]byte, error) {
	if len(command) != 2 {
		return common.RespError("missing key"), fmt.Errorf("missing key")
	}
	var buffer bytes.Buffer
	value, err := GetValueFromMemory(command[1])
	if err != nil {
		return nil, err
	}
	buffer.WriteString("+" + value)
	buffer.WriteString(common.Terminator)
	return []byte(buffer.String()), nil
}

func GetValueFromMemory(key string) (string, error) {
	data, ok := db.DB.Load(key)
	// Todo: Check what redis returns when key not found
	if !ok {
		return "", nil
	}
	return data.(string), nil
}
