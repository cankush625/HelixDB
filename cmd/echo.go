package cmd

import (
	"HelixDB/common"
	"errors"
)

var MessageRequiredError = errors.New("message required")

// Echo returns the message passed to it. Message is required.
func Echo(command common.Cmd) ([]byte, error) {
	if len(command.Args) != 1 {
		return common.RespError("message is required"), MessageRequiredError
	}
	return common.RespBulkString(command.Args[0]), nil
}
