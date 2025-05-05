package cmd

import (
	"errors"
	"reflect"
	"testing"
)

// TestPing tests the Ping function for
// all possible valid and invalid inputs
func TestPing(t *testing.T) {
	tests := []struct {
		command []string
		want    []byte
		wantErr error
	}{
		{[]string{"PING"}, []byte("+PONG\r\n"), nil},
		{[]string{"PING", "Hello"}, []byte("+Hello\r\n"), nil},
		{[]string{"PING", "hello"}, []byte("+hello\r\n"), nil},
	}
	for _, test := range tests {
		if got, gotErr := Ping(test.command); !reflect.DeepEqual(got, test.want) || !errors.Is(gotErr, test.wantErr) {
			t.Errorf("Ping(%v) = %v, %v", test.command, got, gotErr)
		}
	}
}
