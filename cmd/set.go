package cmd

import (
	"HelixDB/common"
	"HelixDB/db"
	"bytes"
	"errors"
	"strconv"
	"strings"
	"time"
)

var SyntaxError = errors.New("syntax error")
var InvalidExpireTimeError = errors.New("invalid expire time in 'set' command")

// flagArgs are options that stand alone with no following value.
var flagArgs = map[string]bool{
	common.NX:     true,
	common.XX:     true,
	common.KEEPTTL: true,
	common.GET:    true,
}

// valueArgs are options that take a following value.
var valueArgs = map[string]bool{
	common.EX:   true,
	common.PX:   true,
	common.EXAT: true,
	common.PXAT: true,
}

func Set(command common.Cmd) ([]byte, error) {
	if len(command.Args) < 2 {
		err := common.WrongNumberOfArgsError(command.Name)
		return common.RespError(err.Error()), err
	}

	key, value := command.Args[0], command.Args[1]

	argsMap, err := parseSetArgs(command.Args)
	if err != nil {
		return common.RespError(err.Error()), err
	}

	// GET — capture old value before making any changes.
	var oldValue string
	var oldExists bool
	if _, hasGet := argsMap[common.GET]; hasGet {
		oldValue, err = GetValueFromMemory(key)
		oldExists = err == nil
	}

	// NX — only set if key does not exist.
	if _, hasNX := argsMap[common.NX]; hasNX {
		if _, err := GetValueFromMemory(key); err == nil {
			return common.RespNullBulkString(), nil
		}
	}

	// XX — only set if key already exists.
	if _, hasXX := argsMap[common.XX]; hasXX {
		if _, err := GetValueFromMemory(key); err != nil {
			return common.RespNullBulkString(), nil
		}
	}

	err = processSetArgs(key, argsMap)
	if err != nil && errors.Is(err, InvalidExpireTimeError) {
		return common.RespError(InvalidExpireTimeError.Error()), InvalidExpireTimeError
	} else if err != nil {
		return common.RespError("error"), err
	}

	setValueInMemory(key, value)

	// Return old value if GET was specified.
	if _, hasGet := argsMap[common.GET]; hasGet {
		if oldExists {
			return common.RespBulkString(oldValue), nil
		}
		return common.RespNullBulkString(), nil
	}

	var buffer bytes.Buffer
	buffer.WriteString("+" + "OK")
	buffer.WriteString(common.Terminator)
	return []byte(buffer.String()), nil
}

// parseSetArgs parses the optional arguments after key and value.
// Flags (NX, XX, KEEPTTL, GET) consume one token.
// Value options (EX, PX, EXAT, PXAT) consume two tokens.
// Returns a map of option name to its value (true for flags, string for value options).
func parseSetArgs(args []string) (map[string]any, error) {
	if len(args) == 2 {
		return nil, nil
	}
	extraArgs := args[2:]
	result := make(map[string]any)

	for i := 0; i < len(extraArgs); {
		// Uppercase for case-insensitive matching — Redis options are case-insensitive.
		arg := strings.ToUpper(extraArgs[i])
		if flagArgs[arg] {
			if _, exists := result[arg]; exists {
				return nil, SyntaxError
			}
			result[arg] = true
			i++
		} else if valueArgs[arg] {
			if i+1 >= len(extraArgs) {
				return nil, SyntaxError
			}
			if _, exists := result[arg]; exists {
				return nil, SyntaxError
			}
			result[arg] = extraArgs[i+1]
			i += 2
		} else {
			return nil, SyntaxError
		}
	}

	// NX and XX are mutually exclusive.
	_, hasNX := result[common.NX]
	_, hasXX := result[common.XX]
	if hasNX && hasXX {
		return nil, SyntaxError
	}

	// EX, PX, EXAT, PXAT, KEEPTTL are all mutually exclusive with each other.
	ttlCount := 0
	for _, opt := range []string{common.EX, common.PX, common.EXAT, common.PXAT, common.KEEPTTL} {
		if _, ok := result[opt]; ok {
			ttlCount++
		}
	}
	if ttlCount > 1 {
		return nil, SyntaxError
	}

	return result, nil
}

// processSetArgs validates TTL args and writes the expiration time to KeyTTL.
// KEEPTTL leaves the existing TTL untouched.
// If no TTL option is provided, the key is marked as persistent (nil expiration).
func processSetArgs(key string, args map[string]any) error {
	if ttl, ok := args[common.EX]; ok {
		seconds, err := strconv.ParseInt(ttl.(string), 10, 64)
		if err != nil || seconds <= 0 {
			return InvalidExpireTimeError
		}
		ttlInMs := common.SecondsToMilliseconds(seconds)
		db.KeyTTL.Store(key, common.GetCurrentTimeInUnixMilli(ttlInMs))
		return nil
	}
	if ttl, ok := args[common.PX]; ok {
		ms, err := strconv.ParseInt(ttl.(string), 10, 64)
		if err != nil || ms <= 0 {
			return InvalidExpireTimeError
		}
		db.KeyTTL.Store(key, common.GetCurrentTimeInUnixMilli(ms))
		return nil
	}
	if ttl, ok := args[common.EXAT]; ok {
		seconds, err := strconv.ParseInt(ttl.(string), 10, 64)
		if err != nil || seconds <= 0 {
			return InvalidExpireTimeError
		}
		db.KeyTTL.Store(key, seconds*int64(time.Second/time.Millisecond))
		return nil
	}
	if ttl, ok := args[common.PXAT]; ok {
		ms, err := strconv.ParseInt(ttl.(string), 10, 64)
		if err != nil || ms <= 0 {
			return InvalidExpireTimeError
		}
		db.KeyTTL.Store(key, ms)
		return nil
	}
	if _, ok := args[common.KEEPTTL]; ok {
		// Leave existing TTL untouched.
		return nil
	}
	// No TTL option — mark as persistent.
	db.KeyTTL.Store(key, nil)
	return nil
}

// setValueInMemory stores the value in the in-memory DB against the key.
func setValueInMemory(key string, value string) {
	db.DB.Store(key, value)
}
