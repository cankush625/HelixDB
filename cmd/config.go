package cmd

import (
	"HelixDB/common"
	"HelixDB/db"
	"errors"
)

var ErrUnknownSubcommand = errors.New("unknown subcommand")
var ErrUnknownConfigParam = errors.New("unknown config parameter")
var ErrInvalidConfigValue = errors.New("invalid config value")

// ConfigCmd handles the CONFIG command.
// Supported subcommands: GET, SET.
func ConfigCmd(command common.Cmd) ([]byte, error) {
	if len(command.Args) < 1 {
		err := common.WrongNumberOfArgsError(command.Name)
		return common.RespError(err.Error()), err
	}

	subcommand := command.Args[0]

	switch subcommand {
	case "GET":
		return configGet(command)
	case "SET":
		return configSet(command)
	default:
		return common.RespError("unknown subcommand '" + subcommand + "' for 'config' command"), ErrUnknownSubcommand
	}
}

// configGet returns the current value of the requested config parameter.
// Usage: CONFIG GET <parameter>
func configGet(command common.Cmd) ([]byte, error) {
	if len(command.Args) != 2 {
		err := common.WrongNumberOfArgsError(command.Name)
		return common.RespError(err.Error()), err
	}
	param := command.Args[1]
	value, ok := db.ServerConfig.Get(param)
	if !ok {
		return common.RespError("unknown config parameter '"+param+"'"), ErrUnknownConfigParam
	}
	// Return as a two-element array: [parameter, value]
	return common.RespBulkStringArray([]string{param, value}), nil
}

// configSet updates the value of the named config parameter.
// Usage: CONFIG SET <parameter> <value>
func configSet(command common.Cmd) ([]byte, error) {
	if len(command.Args) != 3 {
		err := common.WrongNumberOfArgsError(command.Name)
		return common.RespError(err.Error()), err
	}
	param, value := command.Args[1], command.Args[2]
	if ok := db.ServerConfig.Set(param, value); !ok {
		return common.RespError("invalid value for config parameter '"+param+"'"), ErrInvalidConfigValue
	}
	return []byte("+OK" + common.Terminator), nil
}
