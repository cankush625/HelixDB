package cmd

import (
	"HelixDB/common"
	"HelixDB/db"
	"path"
	"time"
)

// Keys returns all keys in the store that match the given glob-style pattern.
// Expired keys are excluded. Returns a RESP bulk string array.
//
// Supported pattern syntax (via path.Match):
//   - * matches any sequence of characters (not including /)
//   - ? matches any single character (not including /)
//   - [abc] matches any character in the set
//
// Note: keys containing '/' are not fully supported with wildcard patterns.
func Keys(command common.Cmd) ([]byte, error) {
	if len(command.Args) != 1 {
		return common.RespError("wrong number of arguments"), WrongNumberOfArgumentsError
	}
	pattern := command.Args[0]
	// Validate pattern before scanning — path.Match returns ErrBadPattern for invalid syntax.
	if _, err := path.Match(pattern, ""); err != nil {
		return common.RespError("invalid pattern"), SyntaxError
	}

	now := time.Now().UnixMilli()
	var keys []string
	db.DB.Range(func(k, _ any) bool {
		key := k.(string)
		// Skip expired keys
		if expirationTime, ok := db.KeyTTL.Load(key); ok {
			if expirationTime != nil && expirationTime.(int64) < now {
				return true
			}
		}
		matched, _ := path.Match(pattern, key)
		if matched {
			keys = append(keys, key)
		}
		return true
	})
	return common.RespBulkStringArray(keys), nil
}
