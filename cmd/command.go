package cmd

import (
	"HelixDB/common"
	"bytes"
)

func Command(command []string) ([]byte, error) {
	var buffer bytes.Buffer
	buffer.WriteString("+")
	if len(command) > 1 && command[1] == "DOCS" {
		buffer.WriteString("OK")
	}
	buffer.WriteString(common.Terminator)
	return []byte(buffer.String()), nil
}
