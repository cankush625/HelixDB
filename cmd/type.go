package cmd

import (
	"HelixDB/common"
	"HelixDB/db"
	"time"
)

// Type returns the type of the value stored at key.
// Returns "none" if the key does not exist or has expired.
// Possible return values: string, integer, float, boolean, json, none.
func Type(command common.Cmd) ([]byte, error) {
	if len(command.Args) != 1 {
		err := common.WrongNumberOfArgsError(command.Name)
		return common.RespError(err.Error()), err
	}

	key := command.Args[0]

	expirationTime, ok := db.KeyTTL.Load(key)
	if !ok {
		return common.RespSimpleString("none"), nil
	}
	if expirationTime != nil && expirationTime.(int64) < time.Now().UnixMilli() {
		return common.RespSimpleString("none"), nil
	}

	data, ok := db.DB.Load(key)
	if !ok {
		return common.RespSimpleString("none"), nil
	}

	return common.RespSimpleString(string(data.(db.Value).Type)), nil
}
