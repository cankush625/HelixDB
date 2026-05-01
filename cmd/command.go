package cmd

import (
	"HelixDB/common"
	"bytes"
)

func Command(command common.Cmd) ([]byte, error) {
	var buffer bytes.Buffer
	buffer.WriteString("+")
	if len(command.Args) > 0 && command.Args[0] == "DOCS" {
		buffer.WriteString("OK")
	}
	buffer.WriteString(common.Terminator)
	return []byte(buffer.String()), nil
}
