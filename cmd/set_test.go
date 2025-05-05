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
		{[]string{"SET", "tenant", "ACME"}, []byte("+OK\r\n"), nil},
		{[]string{"SET", "tenant"}, []byte("-missing arguments\r\n"), MissingArgumentsError},
		{[]string{"SET"}, []byte("-missing arguments\r\n"), MissingArgumentsError},
		{[]string{"SET", "tenant", "ACME", "asd"}, []byte("-syntax error\r\n"), SyntaxError},
	}
	for _, test := range tests {
		if got, gotErr := Set(test.command); !reflect.DeepEqual(got, test.want) || !errors.Is(gotErr, test.wantErr) {
			t.Errorf("Set(%v) = %v, %v", test.command, got, gotErr)
		}
	}
}
