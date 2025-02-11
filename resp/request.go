package resp

import (
	"HelixDB/exc"
	"fmt"
	"net"
)

func PerformRequest(buffer []byte, conn net.Conn) {
	// Todo: This parse command function will return the command, and it's args
	//  Using this command name and args, we need to execute the command
	commandData, err := ParseCommand(buffer)
	if err != nil {
		err := fmt.Errorf("error executing the command")
		fmt.Println(err.Error())
	}
	reply, err := exc.ExecuteCommand(commandData)
	if err != nil {
		err := fmt.Errorf("error executing the command")
		fmt.Println(err.Error())
	}
	ReplyCommand(reply, conn)
}
