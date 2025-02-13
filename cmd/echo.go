package cmd

import (
	"HelixDB/common"
	"bytes"
)

// Echo is a command that returns the message that
// passed to it. Message is required for this command
func Echo(message string) ([]byte, error) {
	var buffer bytes.Buffer
	if message != "" {
		buffer.WriteString("+" + message)
	} else {
		buffer.WriteString("-" + "message is required")
	}
	buffer.WriteString(common.Terminator)
	return []byte(buffer.String()), nil
}
