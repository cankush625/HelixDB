package cmd

import (
	"HelixDB/common"
	"errors"
	"reflect"
	"testing"
)

// TestSet tests the Set function for
// all possible valid and invalid inputs
func TestSet(t *testing.T) {
	tests := []struct {
		command common.Cmd
		want    []byte
		wantErr error
	}{
		// Basic SET
		{common.Cmd{Name: "SET", Args: []string{"tenant", "ACME"}}, []byte("+OK\r\n"), nil},
		// Missing key/value
		{common.Cmd{Name: "SET", Args: []string{"tenant"}}, []byte("-missing arguments\r\n"), MissingArgumentsError},
		{common.Cmd{Name: "SET"}, []byte("-missing arguments\r\n"), MissingArgumentsError},
		// Unknown option
		{common.Cmd{Name: "SET", Args: []string{"tenant", "ACME", "asd"}}, []byte("-syntax error\r\n"), SyntaxError},
		// Valid EX
		{common.Cmd{Name: "SET", Args: []string{"tenant", "ACME", "EX", "10"}}, []byte("+OK\r\n"), nil},
		// Valid PX
		{common.Cmd{Name: "SET", Args: []string{"tenant", "ACME", "PX", "5000"}}, []byte("+OK\r\n"), nil},
		// EX and PX together — mutually exclusive
		{common.Cmd{Name: "SET", Args: []string{"tenant", "ACME", "EX", "10", "PX", "5000"}}, []byte("-syntax error\r\n"), SyntaxError},
		// EX with non-numeric value
		{common.Cmd{Name: "SET", Args: []string{"tenant", "ACME", "EX", "abc"}}, []byte("-invalid expire time in 'set' command\r\n"), InvalidExpireTimeError},
		// EX with zero — must be positive
		{common.Cmd{Name: "SET", Args: []string{"tenant", "ACME", "EX", "0"}}, []byte("-invalid expire time in 'set' command\r\n"), InvalidExpireTimeError},
		// EX with negative value
		{common.Cmd{Name: "SET", Args: []string{"tenant", "ACME", "EX", "-5"}}, []byte("-invalid expire time in 'set' command\r\n"), InvalidExpireTimeError},
		// PX with non-numeric value
		{common.Cmd{Name: "SET", Args: []string{"tenant", "ACME", "PX", "abc"}}, []byte("-invalid expire time in 'set' command\r\n"), InvalidExpireTimeError},
		// PX with zero — must be positive
		{common.Cmd{Name: "SET", Args: []string{"tenant", "ACME", "PX", "0"}}, []byte("-invalid expire time in 'set' command\r\n"), InvalidExpireTimeError},
		// Odd number of extra args (missing value for option)
		{common.Cmd{Name: "SET", Args: []string{"tenant", "ACME", "EX"}}, []byte("-syntax error\r\n"), SyntaxError},
	}
	for _, test := range tests {
		if got, gotErr := Set(test.command); !reflect.DeepEqual(got, test.want) || !errors.Is(gotErr, test.wantErr) {
			t.Errorf("Set(%v) = %v, %v; want %v, %v", test.command, got, gotErr, test.want, test.wantErr)
		}
	}
}
