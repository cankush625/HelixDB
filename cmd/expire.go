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
// RESP integer reply: :1\r\n or :0\r\n
func Expire(command []string) ([]byte, error) {
	if len(command) != 3 {
		return common.RespError("wrong number of arguments"), WrongNumberOfArgumentsError
	}
	seconds, err := strconv.ParseInt(command[2], 10, 64)
	if err != nil || seconds <= 0 {
		return common.RespError(InvalidExpireTimeForExpireError.Error()), InvalidExpireTimeForExpireError
	}
	key := command[1]
	// Key must already exist in DB.
	if _, ok := db.DB.Load(key); !ok {
		return common.RespInteger(0), nil
	}
	ttlInMs := common.SecondsToMilliseconds(seconds)
	expirationTime := common.GetCurrentTimeInUnixMilli(ttlInMs)
	db.KeyTTL.Store(key, expirationTime)
	return common.RespInteger(1), nil
}
