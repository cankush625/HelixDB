package resp

import (
	"HelixDB/common"
	"HelixDB/exc"
	"net"
)

func PerformRequest(buffer []byte, conn net.Conn) {
	command, err := ParseCommand(buffer)
	if err != nil {
		ReplyCommand(common.RespError("invalid request"), conn)
		return
	}
	if command.Name == "" {
		return
	}
	// ExecuteCommand always returns a valid RESP-encoded reply in the []byte,
	// even on error. Send it directly — the error is only for the Go call stack.
	reply, _ := exc.ExecuteCommand(command)
	ReplyCommand(reply, conn)
}
