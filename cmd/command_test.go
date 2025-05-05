package cmd

import (
	"errors"
	"reflect"
	"testing"
)

// TestCommand tests the Command function for
// all possible valid and invalid inputs
func TestCommand(t *testing.T) {
	tests := []struct {
		command []string
		want    []byte
		wantErr error
	}{
		{[]string{"COMMAND"}, []byte("+\r\n"), nil},
		{[]string{"COMMAND", "DOCS"}, []byte("+OK\r\n"), nil},
		{[]string{"COMMAND", "DOCS", "random_arg"}, []byte("+OK\r\n"), nil},
	}
	for _, test := range tests {
		if got, gotErr := Command(test.command); !reflect.DeepEqual(got, test.want) || !errors.Is(gotErr, test.wantErr) {
			t.Errorf("Command(%v) = %v, %v", test.command, got, gotErr)
		}
	}
}
