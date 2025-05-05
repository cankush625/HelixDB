package cmd

import (
	"errors"
	"reflect"
	"testing"
)

// TestEcho tests the Echo function for
// all possible valid and invalid inputs
func TestEcho(t *testing.T) {
	tests := []struct {
		command []string
		want    []byte
		wantErr error
	}{
		{[]string{"ECHO", "hello"}, []byte("+hello\r\n"), nil},
		{[]string{"ECHO"}, []byte("-message is required\r\n"), MessageRequiredError},
	}
	for _, test := range tests {
		if got, gotErr := Echo(test.command); !reflect.DeepEqual(got, test.want) || !errors.Is(gotErr, test.wantErr) {
			t.Errorf("Echo(%v) = %v, %v", test.command, got, gotErr)
		}
	}
}
