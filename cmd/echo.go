package cmd

import (
	"HelixDB/common"
	"bytes"
	"errors"
)

var MessageRequiredError = errors.New("message required")

// Echo is a command that returns the message that
// passed to it. Message is required for this command
func Echo(command []string) ([]byte, error) {
	if len(command) != 2 {
		return common.RespError("message is required"), MessageRequiredError
	}
	var buffer bytes.Buffer
	buffer.WriteString("+" + command[1])
	buffer.WriteString(common.Terminator)
	return []byte(buffer.String()), nil
}
