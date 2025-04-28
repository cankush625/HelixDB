package cmd

import (
	"HelixDB/common"
	"bytes"
)

// Ping is a command that returns a common response PONG when
// no any message or input is passed to this command.
// If any message is passed to this command then it returns
// that message as output.
func Ping(command []string) ([]byte, error) {
	var buffer bytes.Buffer
	buffer.WriteString("+")
	message := "PONG"
	if len(command) > 1 {
		message = command[1]
	}
	buffer.WriteString(message)
	buffer.WriteString(common.Terminator)
	return []byte(buffer.String()), nil
}
