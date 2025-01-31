package exc

import "HelixDB/cmd"

func ExecuteCommand(command []byte) ([]byte, error) {
	return cmd.Ping(command)
}
