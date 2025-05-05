package cmd

import (
	"errors"
	"reflect"
	"testing"
)

// TestGet tests the Get function for
// all possible valid and invalid inputs
func TestGet(t *testing.T) {
	// Prepare
	// Insert a key in the cache
	_, _ = Set([]string{"SET", "tenant", "ACME"})
	tests := []struct {
		command []string
		want    []byte
		wantErr error
	}{
		// Key exists in the cache
		{[]string{"GET", "tenant"}, []byte("+ACMEs\r\n"), nil},
		// Key doesn't exist in the cache
		{[]string{"GET", "org"}, []byte("$-1\r\n"), nil},
		// Wrong number of arguments to the GET command
		{[]string{"GET"}, []byte("-wrong number of arguments\r\n"), WrongNumberOfArgumentsError},
		{[]string{"GET", "tenant", "random_arg"}, []byte("-wrong number of arguments\r\n"), WrongNumberOfArgumentsError},
	}
	for _, test := range tests {
		if got, gotErr := Get(test.command); !reflect.DeepEqual(got, test.want) || !errors.Is(gotErr, test.wantErr) {
			t.Errorf("Get(%v) = %v, %v", test.command, got, gotErr)
		}
	}
}
