package cmd

import (
	"HelixDB/common"
	"bytes"
)

func Command(message string) ([]byte, error) {
	var buffer bytes.Buffer
	buffer.WriteString("+")
	if message == "DOCS" {
		buffer.WriteString("OK")
	}
	buffer.WriteString(common.Terminator)
	return []byte(buffer.String()), nil
}
