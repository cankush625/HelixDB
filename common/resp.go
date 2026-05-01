package common

import (
	"bytes"
	"strconv"
)

func RespInteger(n int) []byte {
	var buffer bytes.Buffer
	buffer.WriteString(":")
	buffer.WriteString(strconv.Itoa(n))
	buffer.WriteString(Terminator)
	return buffer.Bytes()
}
