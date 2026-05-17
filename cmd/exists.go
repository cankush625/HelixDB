package cmd

import (
	"HelixDB/common"
)

// Exists returns the number of the given keys that exist in the store.
// If the same key is provided multiple times it is counted multiple times.
// Expired keys are treated as non-existent.
func Exists(command common.Cmd) ([]byte, error) {
	if len(command.Args) < 1 {
		err := common.WrongNumberOfArgsError(command.Name)
		return common.RespError(err.Error()), err
	}
	count := 0
	for _, key := range command.Args {
		if _, err := GetValueFromMemory(key); err == nil {
			count++
		}
	}
	return common.RespInteger(count), nil
}
