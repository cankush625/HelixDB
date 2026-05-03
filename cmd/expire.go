package cmd

import (
	"HelixDB/common"
	"HelixDB/db"
	"errors"
	"strconv"
)

var InvalidExpireTimeForExpireError = errors.New("invalid expire time in 'expire' command")

// Expire sets the TTL on an existing key in seconds.
// Returns 1 if the key exists and the TTL was set, 0 if the key does not exist.
func Expire(command common.Cmd) ([]byte, error) {
	if len(command.Args) != 2 {
		err := common.WrongNumberOfArgsError(command.Name)
		return common.RespError(err.Error()), err
	}
	seconds, err := strconv.ParseInt(command.Args[1], 10, 64)
	if err != nil || seconds <= 0 {
		return common.RespError(InvalidExpireTimeForExpireError.Error()), InvalidExpireTimeForExpireError
	}
	key := command.Args[0]
	if _, ok := db.DB.Load(key); !ok {
		return common.RespInteger(0), nil
	}
	ttlInMs := common.SecondsToMilliseconds(seconds)
	expirationTime := common.GetCurrentTimeInUnixMilli(ttlInMs)
	db.KeyTTL.Store(key, expirationTime)
	return common.RespInteger(1), nil
}
