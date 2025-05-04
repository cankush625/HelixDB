package exc

import (
	"HelixDB/cmd"
	"errors"
	"reflect"
	"testing"
)

// TestExecuteCommand tests the ExecuteCommand function for
// all possible valid inputs
func TestExecuteCommand(t *testing.T) {
	tests := []struct {
		command []string
		want    []byte
		wantErr error
	}{
		// COMMAND command
		{[]string{"COMMAND"}, []byte("+\r\n"), nil},
		{[]string{"COMMAND", "DOCS"}, []byte("+OK\r\n"), nil},
		{[]string{"COMMAND", "DOCS", "random_arg"}, []byte("+OK\r\n"), nil},
		// PING command
		{[]string{"PING"}, []byte("+PONG\r\n"), nil},
		{[]string{"PING", "Hello"}, []byte("+Hello\r\n"), nil},
		// ECHO command
		{[]string{"ECHO", "hello"}, []byte("+hello\r\n"), nil},
		{[]string{"ECHO"}, []byte("-message is required\r\n"), cmd.MessageRequiredError},
		// SET command
		{[]string{"SET", "tenant", "ACME"}, []byte("+OK\r\n"), nil},
		{[]string{"SET", "tenant"}, []byte("-missing arguments\r\n"), cmd.MissingArgumentsError},
		{[]string{"SET"}, []byte("-missing arguments\r\n"), cmd.MissingArgumentsError},
		{[]string{"SET", "tenant", "ACME", "asd"}, []byte("-syntax error\r\n"), cmd.SyntaxError},
		// GET command
		{[]string{"GET", "tenant"}, []byte("+ACME\r\n"), nil},
		{[]string{"GET", "org"}, []byte("$-1\r\n"), nil},
		{[]string{"GET"}, []byte("-wrong number of arguments\r\n"), cmd.WrongNumberOfArgumentsError},
		{[]string{"GET", "tenant", "random_arg"}, []byte("-wrong number of arguments\r\n"), cmd.WrongNumberOfArgumentsError},
		// Unsupported/Invalid command
	}
	for _, test := range tests {
		if got, gotErr := ExecuteCommand(test.command); !reflect.DeepEqual(got, test.want) || !errors.Is(gotErr, test.wantErr) {
			t.Errorf("ExecuteCommand(%v) = %v, %v", test.command, got, gotErr)
		}
	}
}
