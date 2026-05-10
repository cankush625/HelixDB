package cmd

import (
	"HelixDB/common"
	"HelixDB/db"
	"errors"
	"reflect"
	"testing"
)

// TestType tests the Type command for all possible inputs.
func TestType(t *testing.T) {
	// Seed keys of each type
	db.DB.Store("type_string", db.NewValue("hello"))
	db.KeyTTL.Store("type_string", nil)

	db.DB.Store("type_integer", db.NewValue("42"))
	db.KeyTTL.Store("type_integer", nil)

	db.DB.Store("type_float", db.NewValue("3.14"))
	db.KeyTTL.Store("type_float", nil)

	db.DB.Store("type_boolean", db.NewValue("true"))
	db.KeyTTL.Store("type_boolean", nil)

	db.DB.Store("type_json", db.NewValue(`{"a":1}`))
	db.KeyTTL.Store("type_json", nil)

	tests := []struct {
		command common.Cmd
		want    []byte
		wantErr error
	}{
		{common.Cmd{Name: "TYPE", Args: []string{"type_string"}}, common.RespSimpleString("string"), nil},
		{common.Cmd{Name: "TYPE", Args: []string{"type_integer"}}, common.RespSimpleString("integer"), nil},
		{common.Cmd{Name: "TYPE", Args: []string{"type_float"}}, common.RespSimpleString("float"), nil},
		{common.Cmd{Name: "TYPE", Args: []string{"type_boolean"}}, common.RespSimpleString("boolean"), nil},
		{common.Cmd{Name: "TYPE", Args: []string{"type_json"}}, common.RespSimpleString("json"), nil},
		// Non-existent key
		{common.Cmd{Name: "TYPE", Args: []string{"no_such_key"}}, common.RespSimpleString("none"), nil},
		// Wrong number of arguments
		{common.Cmd{Name: "TYPE"}, common.RespError("wrong number of arguments for 'type' command"), common.ErrWrongNumberOfArgs},
		{common.Cmd{Name: "TYPE", Args: []string{"a", "b"}}, common.RespError("wrong number of arguments for 'type' command"), common.ErrWrongNumberOfArgs},
	}
	for _, test := range tests {
		got, gotErr := Type(test.command)
		if !reflect.DeepEqual(got, test.want) || !errors.Is(gotErr, test.wantErr) {
			t.Errorf("Type(%v) = %v, %v; want %v, %v", test.command, got, gotErr, test.want, test.wantErr)
		}
	}
}
