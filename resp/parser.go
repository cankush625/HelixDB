package resp

import (
	"HelixDB/common"
	"errors"
	"fmt"
	"strings"
)

var UnsupportedCommandDataTypeError = errors.New("unsupported command datatype")

// ParseCommand parses the raw RESP bytes and returns a Cmd.
// Only Array type commands are accepted; Simple Strings are for replies only.
func ParseCommand(command []byte) (common.Cmd, error) {
	if command == nil {
		return common.Cmd{}, nil
	}
	if !isValidCommand(command) {
		return common.Cmd{}, fmt.Errorf("invalid command")
	}
	firstByte := string(command[0])
	dataType, ok := DataTypeToFirstByteMap[firstByte]
	if !ok {
		return common.Cmd{}, UnsupportedCommandDataTypeError
	}
	data := strings.Split(string(command), common.Terminator)
	if dataType == common.Array {
		return parseArray(data)
	}
	return common.Cmd{}, nil
}

func isValidCommand(command []byte) bool {
	return command != nil
}

func parseArray(data []string) (common.Cmd, error) {
	var parts []string
	for i := 2; i < len(data); i += 2 {
		parts = append(parts, data[i])
	}
	if len(parts) == 0 {
		return common.Cmd{}, fmt.Errorf("empty command")
	}
	return common.Cmd{
		Name: strings.ToUpper(parts[0]),
		Args: parts[1:],
	}, nil
}
