package resp

import (
	"HelixDB/common"
	"errors"
	"fmt"
	"reflect"
	"testing"
)

// TestParseCommand tests the ParseCommand function for
// all possible valid and invalid inputs
func TestParseCommand(t *testing.T) {
	tests := []struct {
		command []byte
		want    common.Cmd
		wantErr error
	}{
		// Simple String cannot be accepted as a command — returns zero-value Cmd
		{[]byte("+OK\r\n"), common.Cmd{}, nil},
		// Array type commands
		{[]byte("*1\r\n$4\r\nPING\r\n"), common.Cmd{Name: "PING", Args: []string{}}, nil},
		{[]byte("*1\r\n$4\r\nECHO\r\n"), common.Cmd{Name: "ECHO", Args: []string{}}, nil},
		{[]byte("*1\r\n$4\r\nECHO\r\n$2\r\nhi\r\n"), common.Cmd{Name: "ECHO", Args: []string{"hi"}}, nil},
		// Array type command with multiple args
		{[]byte("*1\r\n$3\r\nSET\r\n$3\r\nkey\r\n$5\r\nvalue\r\n"), common.Cmd{Name: "SET", Args: []string{"key", "value"}}, nil},
		// Invalid commands
		{nil, common.Cmd{}, nil},
		// Unsupported datatype
		{[]byte("%1\r\n$4\r\nECHO\r\n$2\r\nhi\r\n"), common.Cmd{}, UnsupportedCommandDataTypeError},
	}
	for _, test := range tests {
		if got, gotErr := ParseCommand(test.command); !reflect.DeepEqual(got, test.want) || !errors.Is(gotErr, test.wantErr) {
			t.Errorf("ParseCommand(%v) = %v, %v; want %v, %v", test.command, got, gotErr, test.want, test.wantErr)
		}
	}
}

// BenchmarkParseCommand benchmarks the ParseCommand function
// against huge number of executions
func BenchmarkParseCommand(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, err := ParseCommand([]byte("*1\r\n$4\r\nPING\r\n"))
		if err != nil {
			continue
		}
	}
}

// ExampleParseCommand is an example function. It serves as a documentation function.
func ExampleParseCommand() {
	fmt.Println(ParseCommand([]byte("*1\r\n$4\r\nPING\r\n")))
	fmt.Println(ParseCommand([]byte("*1\r\n$4\r\nECHO\r\n$2\r\nhi\r\n")))
	// Output:
	// {PING []} <nil>
	// {ECHO [hi]} <nil>
}
