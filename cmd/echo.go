package cmd

import (
	"HelixDB/common"
	"bytes"
	"errors"
)

var MessageRequiredError = errors.New("message required")

// Echo returns the message passed to it. Message is required.
func Echo(command common.Cmd) ([]byte, error) {
	if len(command.Args) != 1 {
		return common.RespError("message is required"), MessageRequiredError
	}
	var buffer bytes.Buffer
	buffer.WriteString("+" + command.Args[0])
	buffer.WriteString(common.Terminator)
	return []byte(buffer.String()), nil
}
