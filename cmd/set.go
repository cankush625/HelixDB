package cmd

import (
	"HelixDB/common"
	"HelixDB/db"
	"bytes"
	"errors"
	"strconv"
)

var MissingArgumentsError = errors.New("missing arguments")
var SyntaxError = errors.New("syntax error")
var InvalidExpireTimeError = errors.New("invalid expire time in 'set' command")

var SupportedArgs = map[string]bool{
	common.EX: true,
	common.PX: true,
}

func Set(command common.Cmd) ([]byte, error) {
	if len(command.Args) < 2 {
		return common.RespError("missing arguments"), MissingArgumentsError
	}
	argsMap, err := parseSetArgs(command.Args)
	if err != nil && errors.Is(err, SyntaxError) {
		return common.RespError("syntax error"), SyntaxError
	} else if err != nil {
		return common.RespError("error"), err
	}
	err = processSetArgs(command.Args[0], argsMap)
	if err != nil && errors.Is(err, InvalidExpireTimeError) {
		return common.RespError(InvalidExpireTimeError.Error()), InvalidExpireTimeError
	} else if err != nil {
		return common.RespError("error"), err
	}
	_, err = setValueInMemory(command.Args[0], command.Args[1])
	if err != nil {
		return common.RespError("error"), err
	}
	var buffer bytes.Buffer
	buffer.WriteString("+" + "OK")
	buffer.WriteString(common.Terminator)
	return []byte(buffer.String()), nil
}

// parseSetArgs parses the optional arguments after key and value and returns
// a map of argument name to its value.
// EX and PX are mutually exclusive — providing both is a syntax error.
func parseSetArgs(args []string) (map[string]any, error) {
	if len(args) == 2 {
		return nil, nil
	}
	extraArgs := args[2:]
	// Options come in key-value pairs, so extra args must be even in count.
	if len(extraArgs)%2 != 0 {
		return nil, SyntaxError
	}
	result := make(map[string]any, len(extraArgs)/2)
	for i := 0; i < len(extraArgs); i += 2 {
		arg := extraArgs[i]
		if _, ok := SupportedArgs[arg]; !ok {
			return nil, SyntaxError
		}
		if _, exists := result[arg]; exists {
			return nil, SyntaxError
		}
		result[arg] = extraArgs[i+1]
	}
	// EX and PX are mutually exclusive per RESP spec.
	_, hasEX := result[common.EX]
	_, hasPX := result[common.PX]
	if hasEX && hasPX {
		return nil, SyntaxError
	}
	return result, nil
}

// processSetArgs validates TTL args and stores the expiration time in KeyTTL.
// If no TTL is provided, the key is stored as persistent (nil expiration).
func processSetArgs(key string, args map[string]any) error {
	if ttl, ok := args[common.EX]; ok {
		seconds, err := strconv.ParseInt(ttl.(string), 10, 64)
		if err != nil || seconds <= 0 {
			return InvalidExpireTimeError
		}
		ttlInMs := common.SecondsToMilliseconds(seconds)
		expirationTime := common.GetCurrentTimeInUnixMilli(ttlInMs)
		db.KeyTTL.Store(key, expirationTime)
		return nil
	}
	if ttl, ok := args[common.PX]; ok {
		ms, err := strconv.ParseInt(ttl.(string), 10, 64)
		if err != nil || ms <= 0 {
			return InvalidExpireTimeError
		}
		expirationTime := common.GetCurrentTimeInUnixMilli(ms)
		db.KeyTTL.Store(key, expirationTime)
		return nil
	}
	// No TTL provided — mark as persistent.
	db.KeyTTL.Store(key, nil)
	return nil
}

// setValueInMemory stores the value in the in-memory DB against the key.
func setValueInMemory(key string, value string) (bool, error) {
	db.DB.Store(key, value)
	return true, nil
}
