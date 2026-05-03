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
		err := common.WrongNumberOfArgsError(command.Name)
		return common.RespError(err.Error()), err
	}
	pattern := command.Args[0]
	// Validate pattern before scanning — path.Match returns ErrBadPattern for invalid syntax.
	if _, err := path.Match(pattern, ""); err != nil {
		return common.RespError("invalid pattern"), SyntaxError
	}

	now := time.Now().UnixMilli()
	matchAll := pattern == "*"
	var keys []string

	// Iterate KeyTTL as the single source of truth for all live keys,
	// avoiding a separate DB lookup per key.
	db.KeyTTL.Range(func(k, value any) bool {
		// Skip expired keys
		if value != nil && value.(int64) < now {
			return true
		}
		key := k.(string)
		if matchAll {
			keys = append(keys, key)
			return true
		}
		if matched, _ := path.Match(pattern, key); matched {
			keys = append(keys, key)
		}
		return true
	})
	return common.RespBulkStringArray(keys), nil
}
