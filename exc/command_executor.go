package exc

import (
	"HelixDB/cmd"
	"HelixDB/common"
	"errors"
)

var UnsupportedCommand = errors.New("unsupported command")

func ExecuteCommand(command common.Cmd) ([]byte, error) {
	switch command.Name {
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
	case common.Expire:
		return cmd.Expire(command)
	case common.Del:
		return cmd.Del(command)
	case common.Keys:
		return cmd.Keys(command)
	case common.Config:
		return cmd.ConfigCmd(command)
	case common.Type:
		return cmd.Type(command)
	}
	return []byte("-unsupported command\r\n"), UnsupportedCommand
}
