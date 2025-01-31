package resp

import (
	"HelixDB/exc"
	"fmt"
	"net"
)

func PerformRequest(buffer []byte, conn net.Conn) {
	ParseCommand(buffer, conn)
	reply, err := exc.ExecuteCommand(buffer)
	if err != nil {
		err := fmt.Errorf("error executing the command")
		fmt.Println(err.Error())
	}
	ReplyCommand(reply, conn)
}
