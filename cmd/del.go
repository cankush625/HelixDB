package cmd

import (
	"HelixDB/common"
	"HelixDB/db"
)

// Del deletes one or more keys from the store.
// Keys that do not exist are ignored and not counted.
// Returns the number of keys that were actually deleted as a RESP integer.
func Del(command []string) ([]byte, error) {
	if len(command) < 2 {
		return common.RespError("wrong number of arguments"), WrongNumberOfArgumentsError
	}
	deleted := 0
	for _, key := range command[1:] {
		if _, ok := db.DB.Load(key); ok {
			db.DB.Delete(key)
			db.KeyTTL.Delete(key)
			deleted++
		}
	}
	return common.RespInteger(deleted), nil
}
