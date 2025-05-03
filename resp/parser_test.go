package resp

import (
	"errors"
	"fmt"
	"reflect"
	"testing"
)

// TestParseCommand tests the ParseCommand function for
// all possible valid inputs
func TestParseCommand(t *testing.T) {
	tests := []struct {
		command []byte
		want    []string
		wantErr error
	}{
		// Simple String cannot be accepted as a command.
		{[]byte("+OK\r\n"), nil, nil},
		// Array type command
		{[]byte("*1\r\n$4\r\nPING\r\n"), []string{"PING"}, nil},
		{[]byte("*1\r\n$4\r\nECHO\r\n"), []string{"ECHO"}, nil},
		{[]byte("*1\r\n$4\r\nECHO\r\n$2\r\nhi\r\n"), []string{"ECHO", "hi"}, nil},
		// Invalid commands
		{nil, nil, nil},
		// Unsupported datatype
		{[]byte("%1\r\n$4\r\nECHO\r\n$2\r\nhi\r\n"), nil, UnsupportedCommandError},
	}
	for _, test := range tests {
		if got, gotErr := ParseCommand(test.command); !reflect.DeepEqual(got, test.want) || !errors.Is(gotErr, test.wantErr) {
			fmt.Println(reflect.TypeOf(gotErr))
			t.Errorf("ParseCommand(%v) = %v, %v", test.command, got, gotErr)
		}
	}
}

// BenchmarkGetDivision benchmarks the ParseCommand function
// against huge number of executions
func BenchmarkParseCommand(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, err := ParseCommand([]byte("*1\r\n$4\r\nPING\r\n"))
		if err != nil {
			continue
		}
	}
}

// ExampleParseCommand is an example function. It serves as a
// documentation function.
func ExampleParseCommand() {
	fmt.Println(ParseCommand([]byte("*1\r\n$4\r\nPING\r\n")))
	fmt.Println(ParseCommand([]byte("*1\r\n$4\r\nECHO\r\n$2\r\nhi\r\n")))
	// Output:
	// [PING] <nil>
	// [ECHO hi] <nil>
}
