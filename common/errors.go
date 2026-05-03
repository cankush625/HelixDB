package common

import (
	"bytes"
	"errors"
	"fmt"
	"strings"
)

func RespError(message string) []byte {
	var buffer bytes.Buffer
	buffer.WriteString("-" + message)
	buffer.WriteString(Terminator)
	return []byte(buffer.String())
}

// ErrWrongNumberOfArgs is the sentinel for wrong-argument-count errors.
// Use errors.Is(err, common.ErrWrongNumberOfArgs) to check.
var ErrWrongNumberOfArgs = errors.New("wrong number of arguments")

// wrongNumberOfArgsError is a command-specific error that satisfies errors.Is(err, ErrWrongNumberOfArgs).
type wrongNumberOfArgsError struct {
	cmdName string
}

func (e *wrongNumberOfArgsError) Error() string {
	return fmt.Sprintf("wrong number of arguments for '%s' command", strings.ToLower(e.cmdName))
}

func (e *wrongNumberOfArgsError) Is(target error) bool {
	return target == ErrWrongNumberOfArgs
}

// WrongNumberOfArgsError returns a command-specific error wrapping ErrWrongNumberOfArgs.
func WrongNumberOfArgsError(cmdName string) error {
	return &wrongNumberOfArgsError{cmdName: cmdName}
}
