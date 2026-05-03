package cmd

import (
	"HelixDB/common"
	"errors"
	"reflect"
	"testing"
)

func TestDel(t *testing.T) {
	// Prepare: insert keys
	_, _ = Set(common.Cmd{Name: "SET", Args: []string{"key1", "value1"}})
	_, _ = Set(common.Cmd{Name: "SET", Args: []string{"key2", "value2"}})
	_, _ = Set(common.Cmd{Name: "SET", Args: []string{"key3", "value3"}})

	tests := []struct {
		command common.Cmd
		want    []byte
		wantErr error
	}{
		// Delete a single existing key — returns 1
		{common.Cmd{Name: "DEL", Args: []string{"key1"}}, []byte(":1\r\n"), nil},
		// Delete a key that does not exist — returns 0
		{common.Cmd{Name: "DEL", Args: []string{"ghost"}}, []byte(":0\r\n"), nil},
		// Delete multiple keys — key2 exists, key3 exists, ghost does not
		{common.Cmd{Name: "DEL", Args: []string{"key2", "key3", "ghost"}}, []byte(":2\r\n"), nil},
		// Wrong number of arguments
		{common.Cmd{Name: "DEL"}, []byte("-wrong number of arguments for 'del' command\r\n"), common.ErrWrongNumberOfArgs},
	}
	for _, test := range tests {
		if got, gotErr := Del(test.command); !reflect.DeepEqual(got, test.want) || !errors.Is(gotErr, test.wantErr) {
			t.Errorf("Del(%v) = %v, %v; want %v, %v", test.command, got, gotErr, test.want, test.wantErr)
		}
	}
}
