package cmd

import (
	"HelixDB/common"
	"errors"
	"reflect"
	"testing"
)

func TestExists(t *testing.T) {
	_, _ = Set(common.Cmd{Name: "SET", Args: []string{"ex_key1", "value1"}})
	_, _ = Set(common.Cmd{Name: "SET", Args: []string{"ex_key2", "value2"}})

	tests := []struct {
		command common.Cmd
		want    []byte
		wantErr error
	}{
		// Single existing key
		{common.Cmd{Name: "EXISTS", Args: []string{"ex_key1"}}, common.RespInteger(1), nil},
		// Single non-existent key
		{common.Cmd{Name: "EXISTS", Args: []string{"no_such_key"}}, common.RespInteger(0), nil},
		// Multiple keys, all exist
		{common.Cmd{Name: "EXISTS", Args: []string{"ex_key1", "ex_key2"}}, common.RespInteger(2), nil},
		// Multiple keys, some missing
		{common.Cmd{Name: "EXISTS", Args: []string{"ex_key1", "no_such_key"}}, common.RespInteger(1), nil},
		// Duplicate key counts multiple times
		{common.Cmd{Name: "EXISTS", Args: []string{"ex_key1", "ex_key1"}}, common.RespInteger(2), nil},
		// Wrong number of arguments
		{common.Cmd{Name: "EXISTS"}, common.RespError("wrong number of arguments for 'exists' command"), common.ErrWrongNumberOfArgs},
	}
	for _, test := range tests {
		if got, gotErr := Exists(test.command); !reflect.DeepEqual(got, test.want) || !errors.Is(gotErr, test.wantErr) {
			t.Errorf("Exists(%v) = %v, %v; want %v, %v", test.command, got, gotErr, test.want, test.wantErr)
		}
	}
}
