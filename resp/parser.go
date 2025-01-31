package resp

import (
	"fmt"
	"net"
)

func ParseCommand(command []byte, conn net.Conn) []byte {
	if command != nil {
		fmt.Printf("parse command: %s", command)
		return command
	}
	return nil
}
