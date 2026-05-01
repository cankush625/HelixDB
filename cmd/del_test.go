package cmd

import (
	"errors"
	"reflect"
	"testing"
)

func TestDel(t *testing.T) {
	// Prepare: insert keys
	_, _ = Set([]string{"SET", "key1", "value1"})
	_, _ = Set([]string{"SET", "key2", "value2"})
	_, _ = Set([]string{"SET", "key3", "value3"})

	tests := []struct {
		command []string
		want    []byte
		wantErr error
	}{
		// Delete a single existing key — returns 1
		{[]string{"DEL", "key1"}, []byte(":1\r\n"), nil},
		// Delete a key that does not exist — returns 0
		{[]string{"DEL", "ghost"}, []byte(":0\r\n"), nil},
		// Delete multiple keys — key2 exists, key3 exists, ghost does not
		{[]string{"DEL", "key2", "key3", "ghost"}, []byte(":2\r\n"), nil},
		// Wrong number of arguments
		{[]string{"DEL"}, []byte("-wrong number of arguments\r\n"), WrongNumberOfArgumentsError},
	}
	for _, test := range tests {
		if got, gotErr := Del(test.command); !reflect.DeepEqual(got, test.want) || !errors.Is(gotErr, test.wantErr) {
			t.Errorf("Del(%v) = %v, %v; want %v, %v", test.command, got, gotErr, test.want, test.wantErr)
		}
	}
}
