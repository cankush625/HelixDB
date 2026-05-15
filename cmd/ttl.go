package cmd

import (
	"HelixDB/common"
	"HelixDB/db"
	"time"
)

// TTL returns the remaining time-to-live of a key in seconds.
// Returns -1 if the key exists but has no TTL.
// Returns -2 if the key does not exist or has expired.
func TTL(command common.Cmd) ([]byte, error) {
	if len(command.Args) != 1 {
		err := common.WrongNumberOfArgsError(command.Name)
		return common.RespError(err.Error()), err
	}
	return ttlMillis(command.Args[0], false), nil
}

// PTTL returns the remaining time-to-live of a key in milliseconds.
// Returns -1 if the key exists but has no TTL.
// Returns -2 if the key does not exist or has expired.
func PTTL(command common.Cmd) ([]byte, error) {
	if len(command.Args) != 1 {
		err := common.WrongNumberOfArgsError(command.Name)
		return common.RespError(err.Error()), err
	}
	return ttlMillis(command.Args[0], true), nil
}

// ttlMillis looks up the TTL for key and returns it as a RESP integer.
// If millis is true the value is in milliseconds, otherwise seconds.
func ttlMillis(key string, millis bool) []byte {
	expiry, ok := db.KeyTTL.Load(key)
	if !ok {
		// Key does not exist at all.
		return common.RespInteger(-2)
	}
	if expiry == nil {
		// Key exists with no TTL.
		return common.RespInteger(-1)
	}
	remaining := expiry.(int64) - time.Now().UnixMilli()
	if remaining <= 0 {
		// Key has expired (passive check).
		return common.RespInteger(-2)
	}
	if millis {
		return common.RespInteger(int(remaining))
	}
	return common.RespInteger(int(remaining / 1000))
}
