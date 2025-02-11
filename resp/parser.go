package resp

import (
	"HelixDB/common"
	"fmt"
	"strings"
)

func ParseCommand(command []byte) (map[string]string, error) {
	if command == nil {
		return nil, nil
	}
	if !isValidCommand(command) {
		return nil, fmt.Errorf("invalid command")
	}
	firstByte := string(command[0])
	dataType, ok := DataTypes[firstByte]
	if !ok {
		return nil, fmt.Errorf("unsupported command")
	}
	data := strings.Split(string(command), common.Terminator)
	if dataType == Array {
		return parseArray(data)
	}
	return nil, nil
}

func isValidCommand(command []byte) bool {
	if command == nil {
		return false
	}
	return true
}

func parseArray(data []string) (map[string]string, error) {
	var command string
	var value string
	for i, datum := range data {
		if i == 2 {
			command = datum
		}
		if i == 4 {
			value = datum
		}
	}
	return map[string]string{
		"command": command,
		"value":   value,
	}, nil
}
