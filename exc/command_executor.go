package exc

import (
	"HelixDB/cmd"
	"HelixDB/common"
	"errors"
	"strings"
)

var UnsupportedCommand = errors.New("unsupported command")

func ExecuteCommand(command []string) ([]byte, error) {
	switch strings.ToUpper(command[0]) {
	case common.Ping:
		return cmd.Ping(command)
	case common.Command:
		return cmd.Command(command)
	case common.Echo:
		return cmd.Echo(command)
	case common.Get:
		return cmd.Get(command)
	case common.Set:
		return cmd.Set(command)
	}
	return []byte("-unsupported command\r\n"), UnsupportedCommand
}
