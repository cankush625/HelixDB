package exc

import (
	"HelixDB/cmd"
	"HelixDB/common"
	"fmt"
	"strings"
)

func ExecuteCommand(command map[string]string) ([]byte, error) {
	switch strings.ToUpper(command["command"]) {
	case common.Ping:
		return cmd.Ping(command["value"])
	case common.Command:
		return cmd.Command(command["value"])
	case common.Echo:
		return cmd.Echo(command["value"])
	}
	return []byte("-unsupported command\r\n"), fmt.Errorf("unsupported command")
}
