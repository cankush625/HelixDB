package cmd

import (
	"errors"
	"reflect"
	"testing"
)

// TestSet tests the Set function for
// all possible valid and invalid inputs
func TestSet(t *testing.T) {
	tests := []struct {
		command []string
		want    []byte
		wantErr error
	}{
		// Basic SET
		{[]string{"SET", "tenant", "ACME"}, []byte("+OK\r\n"), nil},
		// Missing key/value
		{[]string{"SET", "tenant"}, []byte("-missing arguments\r\n"), MissingArgumentsError},
		{[]string{"SET"}, []byte("-missing arguments\r\n"), MissingArgumentsError},
		// Unknown option
		{[]string{"SET", "tenant", "ACME", "asd"}, []byte("-syntax error\r\n"), SyntaxError},
		// Valid EX
		{[]string{"SET", "tenant", "ACME", "EX", "10"}, []byte("+OK\r\n"), nil},
		// Valid PX
		{[]string{"SET", "tenant", "ACME", "PX", "5000"}, []byte("+OK\r\n"), nil},
		// EX and PX together — mutually exclusive
		{[]string{"SET", "tenant", "ACME", "EX", "10", "PX", "5000"}, []byte("-syntax error\r\n"), SyntaxError},
		// EX with non-numeric value
		{[]string{"SET", "tenant", "ACME", "EX", "abc"}, []byte("-invalid expire time in 'set' command\r\n"), InvalidExpireTimeError},
		// EX with zero — must be positive
		{[]string{"SET", "tenant", "ACME", "EX", "0"}, []byte("-invalid expire time in 'set' command\r\n"), InvalidExpireTimeError},
		// EX with negative value
		{[]string{"SET", "tenant", "ACME", "EX", "-5"}, []byte("-invalid expire time in 'set' command\r\n"), InvalidExpireTimeError},
		// PX with non-numeric value
		{[]string{"SET", "tenant", "ACME", "PX", "abc"}, []byte("-invalid expire time in 'set' command\r\n"), InvalidExpireTimeError},
		// PX with zero — must be positive
		{[]string{"SET", "tenant", "ACME", "PX", "0"}, []byte("-invalid expire time in 'set' command\r\n"), InvalidExpireTimeError},
		// Odd number of extra args (missing value for option)
		{[]string{"SET", "tenant", "ACME", "EX"}, []byte("-syntax error\r\n"), SyntaxError},
	}
	for _, test := range tests {
		if got, gotErr := Set(test.command); !reflect.DeepEqual(got, test.want) || !errors.Is(gotErr, test.wantErr) {
			t.Errorf("Set(%v) = %v, %v; want %v, %v", test.command, got, gotErr, test.want, test.wantErr)
		}
	}
}
