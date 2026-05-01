package cmd

import (
	"HelixDB/common"
	"bytes"
)

// Ping returns PONG when no argument is provided,
// or echoes back the first argument if one is given.
func Ping(command common.Cmd) ([]byte, error) {
	var buffer bytes.Buffer
	buffer.WriteString("+")
	message := "PONG"
	if len(command.Args) > 0 {
		message = command.Args[0]
	}
	buffer.WriteString(message)
	buffer.WriteString(common.Terminator)
	return []byte(buffer.String()), nil
}
