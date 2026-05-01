package cmd

import (
	"HelixDB/common"
	"errors"
	"reflect"
	"testing"
)

// TestCommand tests the Command function for
// all possible valid and invalid inputs
func TestCommand(t *testing.T) {
	tests := []struct {
		command common.Cmd
		want    []byte
		wantErr error
	}{
		{common.Cmd{Name: "COMMAND"}, []byte("+\r\n"), nil},
		{common.Cmd{Name: "COMMAND", Args: []string{"DOCS"}}, []byte("+OK\r\n"), nil},
		{common.Cmd{Name: "COMMAND", Args: []string{"DOCS", "random_arg"}}, []byte("+OK\r\n"), nil},
	}
	for _, test := range tests {
		if got, gotErr := Command(test.command); !reflect.DeepEqual(got, test.want) || !errors.Is(gotErr, test.wantErr) {
			t.Errorf("Command(%v) = %v, %v; want %v, %v", test.command, got, gotErr, test.want, test.wantErr)
		}
	}
}
