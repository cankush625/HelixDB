package resp

import (
	"fmt"
)

func ParseCommand(command []byte) []byte {
	if command != nil {
		fmt.Printf("parse command: %s", command)
		return command
	}
	return nil
}
