package cmd

import (
	"HelixDB/common"
	"HelixDB/db"
	"bytes"
	"fmt"
)

func Set(command []string) ([]byte, error) {
	if len(command) != 3 {
		return common.RespError("missing arguments"), fmt.Errorf("missing arguments")
	}
	var buffer bytes.Buffer
	_, err := SetValueFromMemory(command[1], command[2])
	if err != nil {
		return nil, err
	}
	buffer.WriteString("+" + "OK")
	buffer.WriteString(common.Terminator)
	return []byte(buffer.String()), nil
}

func SetValueFromMemory(key string, value string) (bool, error) {
	db.DB.Store(key, value)
	return true, nil
}
