package cmd

import (
	"HelixDB/common"
	"errors"
	"reflect"
	"testing"
)

// TestPing tests the Ping function for
// all possible valid and invalid inputs
func TestPing(t *testing.T) {
	tests := []struct {
		command common.Cmd
		want    []byte
		wantErr error
	}{
		{common.Cmd{Name: "PING"}, []byte("+PONG\r\n"), nil},
		{common.Cmd{Name: "PING", Args: []string{"Hello"}}, []byte("+Hello\r\n"), nil},
		{common.Cmd{Name: "PING", Args: []string{"hello"}}, []byte("+hello\r\n"), nil},
	}
	for _, test := range tests {
		if got, gotErr := Ping(test.command); !reflect.DeepEqual(got, test.want) || !errors.Is(gotErr, test.wantErr) {
			t.Errorf("Ping(%v) = %v, %v; want %v, %v", test.command, got, gotErr, test.want, test.wantErr)
		}
	}
}
