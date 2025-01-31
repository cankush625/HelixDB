package resp

import (
	"fmt"
	"net"
)

func ReplyCommand(reply []byte, conn net.Conn) {
	_, err := conn.Write(reply)
	if err != nil {
		err := fmt.Errorf("error writing to a connection")
		fmt.Println(err.Error())
	}
}
