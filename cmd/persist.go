package cmd

import (
	"HelixDB/common"
	"HelixDB/db"
	"time"
)

// Persist removes the TTL from a key so that it never expires.
// Returns 1 if an existing timeout was removed.
// Returns 0 if the key does not exist, has no TTL, or has already expired.
func Persist(command common.Cmd) ([]byte, error) {
	if len(command.Args) != 1 {
		err := common.WrongNumberOfArgsError(command.Name)
		return common.RespError(err.Error()), err
	}
	key := command.Args[0]
	// Nothing to remove when the key is absent (!ok), has no TTL (ttl == nil),
	// or holds a timeout that has already elapsed (treated as gone, like GET/TTL).
	if ttl, ok := db.KeyTTL.Load(key); !ok || ttl == nil || ttl.(int64) < time.Now().UnixMilli() {
		return common.RespInteger(0), nil
	}
	db.KeyTTL.Store(key, nil)
	return common.RespInteger(1), nil
}
