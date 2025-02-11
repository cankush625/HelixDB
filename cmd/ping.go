package cmd

import (
	"HelixDB/common"
	"bytes"
)

// Ping is a command that returns a common response PONG when
// no any message or input is passed to this command.
// If any message is passed to this command then it returns
// that message as output.
func Ping(message string) ([]byte, error) {
	var buffer bytes.Buffer
	buffer.WriteString("+")
	if message != "" {
		buffer.WriteString(message)
	} else {
		buffer.WriteString("PONG")
	}
	buffer.WriteString(common.Terminator)
	return []byte(buffer.String()), nil
}
