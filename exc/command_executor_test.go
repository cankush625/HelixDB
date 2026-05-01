package exc

import (
	"HelixDB/cmd"
	"HelixDB/common"
	"errors"
	"fmt"
	"reflect"
	"testing"
)

// TestExecuteCommand tests the ExecuteCommand function for
// all possible valid and invalid inputs
func TestExecuteCommand(t *testing.T) {
	tests := []struct {
		command common.Cmd
		want    []byte
		wantErr error
	}{
		// COMMAND command
		{common.Cmd{Name: "COMMAND"}, []byte("+\r\n"), nil},
		{common.Cmd{Name: "COMMAND", Args: []string{"DOCS"}}, []byte("+OK\r\n"), nil},
		{common.Cmd{Name: "COMMAND", Args: []string{"DOCS", "random_arg"}}, []byte("+OK\r\n"), nil},
		// PING command
		{common.Cmd{Name: "PING"}, []byte("+PONG\r\n"), nil},
		{common.Cmd{Name: "PING", Args: []string{"Hello"}}, []byte("+Hello\r\n"), nil},
		// ECHO command
		{common.Cmd{Name: "ECHO", Args: []string{"hello"}}, []byte("+hello\r\n"), nil},
		{common.Cmd{Name: "ECHO"}, []byte("-message is required\r\n"), cmd.MessageRequiredError},
		// SET command
		{common.Cmd{Name: "SET", Args: []string{"tenant", "ACME"}}, []byte("+OK\r\n"), nil},
		{common.Cmd{Name: "SET", Args: []string{"tenant"}}, []byte("-missing arguments\r\n"), cmd.MissingArgumentsError},
		{common.Cmd{Name: "SET"}, []byte("-missing arguments\r\n"), cmd.MissingArgumentsError},
		{common.Cmd{Name: "SET", Args: []string{"tenant", "ACME", "asd"}}, []byte("-syntax error\r\n"), cmd.SyntaxError},
		// GET command
		{common.Cmd{Name: "GET", Args: []string{"tenant"}}, []byte("+ACME\r\n"), nil},
		{common.Cmd{Name: "GET", Args: []string{"org"}}, []byte("$-1\r\n"), nil},
		{common.Cmd{Name: "GET"}, []byte("-wrong number of arguments\r\n"), cmd.WrongNumberOfArgumentsError},
		{common.Cmd{Name: "GET", Args: []string{"tenant", "random_arg"}}, []byte("-wrong number of arguments\r\n"), cmd.WrongNumberOfArgumentsError},
		// Unsupported/Invalid command
		{common.Cmd{Name: "UNSUPPORTED"}, []byte("-unsupported command\r\n"), UnsupportedCommand},
	}
	for _, test := range tests {
		if got, gotErr := ExecuteCommand(test.command); !reflect.DeepEqual(got, test.want) || !errors.Is(gotErr, test.wantErr) {
			t.Errorf("ExecuteCommand(%v) = %v, %v; want %v, %v", test.command, got, gotErr, test.want, test.wantErr)
		}
	}
}

// BenchmarkExecuteCommand benchmarks the ExecuteCommand function
// against huge number of executions
func BenchmarkExecuteCommand(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, err := ExecuteCommand(common.Cmd{Name: "PING"})
		if err != nil {
			continue
		}
	}
}

// ExampleExecuteCommand is an example function. It serves as a documentation function.
func ExampleExecuteCommand() {
	fmt.Println(ExecuteCommand(common.Cmd{Name: "PING"}))
	fmt.Println(ExecuteCommand(common.Cmd{Name: "ECHO", Args: []string{"hello"}}))
	fmt.Println(ExecuteCommand(common.Cmd{Name: "SET", Args: []string{"tenant", "ACME"}}))
	fmt.Println(ExecuteCommand(common.Cmd{Name: "GET", Args: []string{"tenant"}}))
	// Output:
	// [43 80 79 78 71 13 10] <nil>
	// [43 104 101 108 108 111 13 10] <nil>
	// [43 79 75 13 10] <nil>
	// [43 65 67 77 69 13 10] <nil>
}
