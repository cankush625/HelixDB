package cmd

import (
	"HelixDB/common"
	"errors"
	"reflect"
	"testing"
)

// TestEcho tests the Echo function for
// all possible valid and invalid inputs
func TestEcho(t *testing.T) {
	tests := []struct {
		command common.Cmd
		want    []byte
		wantErr error
	}{
		{common.Cmd{Name: "ECHO", Args: []string{"hello"}}, []byte("+hello\r\n"), nil},
		{common.Cmd{Name: "ECHO"}, []byte("-message is required\r\n"), MessageRequiredError},
	}
	for _, test := range tests {
		if got, gotErr := Echo(test.command); !reflect.DeepEqual(got, test.want) || !errors.Is(gotErr, test.wantErr) {
			t.Errorf("Echo(%v) = %v, %v; want %v, %v", test.command, got, gotErr, test.want, test.wantErr)
		}
	}
}
