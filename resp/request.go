package resp

import (
	"HelixDB/exc"
	"net"
)

func PerformRequest(buffer []byte, conn net.Conn) {
	// Todo: This parse command function will return the command, and it's args
	//  Using this command name and args, we need to execute the command
	commandData, err := ParseCommand(buffer)
	if err != nil {
		ReplyCommand([]byte("-unsupported command\r\n"), conn)
		return
	}
	reply, err := exc.ExecuteCommand(commandData)
	if err != nil {
		ReplyCommand([]byte("-error executing the command\r\n"), conn)
		return
	}
	ReplyCommand(reply, conn)
}
