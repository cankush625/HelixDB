package resp

import (
	"HelixDB/common"
	"errors"
	"fmt"
	"strings"
)

var UnsupportedCommandError = errors.New("unsupported command")

// ParseCommand parses the command and returns it as a slice of string.
// ParseCommand parses the commands of type Array.
//
// Note:
// Simple String cannot be accepted as a command and is only for command replies.
func ParseCommand(command []byte) ([]string, error) {
	if command == nil {
		return nil, nil
	}
	if !isValidCommand(command) {
		return nil, fmt.Errorf("invalid command")
	}
	firstByte := string(command[0])
	dataType, ok := DataTypes[firstByte]
	if !ok {
		return nil, UnsupportedCommandError
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

func parseArray(data []string) ([]string, error) {
	var value []string
	for i := 2; i < len(data); i += 2 {
		value = append(value, data[i])
	}
	return value, nil
}
