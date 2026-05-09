package common

import (
	"bytes"
	"fmt"
	"strconv"
)

// RespBulkString formats a string value as a RESP bulk string: $len\r\nvalue\r\n
func RespBulkString(value string) []byte {
	var buffer bytes.Buffer
	buffer.WriteString(fmt.Sprintf("$%d", len(value)))
	buffer.WriteString(Terminator)
	buffer.WriteString(value)
	buffer.WriteString(Terminator)
	return buffer.Bytes()
}

// RespNullBulkString returns the RESP null bulk string: $-1\r\n
// Used when a key does not exist or has expired.
func RespNullBulkString() []byte {
	return []byte("$-1" + Terminator)
}

// RespBulkStringArray formats a slice of strings as a RESP array of bulk strings.
// An empty slice returns *0\r\n (empty array).
func RespBulkStringArray(items []string) []byte {
	var buffer bytes.Buffer
	buffer.WriteString(fmt.Sprintf("*%d", len(items)))
	buffer.WriteString(Terminator)
	for _, item := range items {
		buffer.WriteString(fmt.Sprintf("$%d", len(item)))
		buffer.WriteString(Terminator)
		buffer.WriteString(item)
		buffer.WriteString(Terminator)
	}
	return buffer.Bytes()
}

// RespSimpleString formats a value as a RESP simple string: +value\r\n
func RespSimpleString(value string) []byte {
	return []byte("+" + value + Terminator)
}

func RespInteger(n int) []byte {
	var buffer bytes.Buffer
	buffer.WriteString(":")
	buffer.WriteString(strconv.Itoa(n))
	buffer.WriteString(Terminator)
	return buffer.Bytes()
}
