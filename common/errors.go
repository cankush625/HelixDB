package common

import (
	"bytes"
)

func RespError(message string) []byte {
	var buffer bytes.Buffer
	buffer.WriteString("-" + message)
	buffer.WriteString(Terminator)
	return []byte(buffer.String())
}
