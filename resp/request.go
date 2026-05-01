package resp

import (
	"HelixDB/exc"
	"net"
)

func PerformRequest(buffer []byte, conn net.Conn) {
	command, err := ParseCommand(buffer)
	if err != nil {
		ReplyCommand([]byte("-unsupported command\r\n"), conn)
		return
	}
	if command.Name == "" {
		return
	}
	reply, err := exc.ExecuteCommand(command)
	if err != nil {
		ReplyCommand([]byte("-error executing the command\r\n"), conn)
		return
	}
	ReplyCommand(reply, conn)
}
