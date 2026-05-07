package cmd

import (
	"HelixDB/common"
	"HelixDB/db"
	"errors"
	"reflect"
	"testing"
)

func TestConfigGet(t *testing.T) {
	// Reset config to defaults before testing
	db.ServerConfig.Set("hz", "1")
	db.ServerConfig.Set("active-expire-enabled", "yes")
	db.ServerConfig.Set("maxmemory", "0")

	tests := []struct {
		command common.Cmd
		want    []byte
		wantErr error
	}{
		// Valid GET
		{common.Cmd{Name: "CONFIG", Args: []string{"GET", "hz"}}, common.RespBulkStringArray([]string{"hz", "1"}), nil},
		{common.Cmd{Name: "CONFIG", Args: []string{"GET", "active-expire-enabled"}}, common.RespBulkStringArray([]string{"active-expire-enabled", "yes"}), nil},
		{common.Cmd{Name: "CONFIG", Args: []string{"GET", "maxmemory"}}, common.RespBulkStringArray([]string{"maxmemory", "0"}), nil},
		// Unknown parameter
		{common.Cmd{Name: "CONFIG", Args: []string{"GET", "unknown"}}, common.RespError("unknown config parameter 'unknown'"), ErrUnknownConfigParam},
		// Wrong number of arguments
		{common.Cmd{Name: "CONFIG", Args: []string{"GET"}}, common.RespError("wrong number of arguments for 'config' command"), common.ErrWrongNumberOfArgs},
		// Case-insensitive subcommand
		{common.Cmd{Name: "CONFIG", Args: []string{"get", "hz"}}, common.RespBulkStringArray([]string{"hz", "1"}), nil},
	}
	for _, test := range tests {
		got, gotErr := ConfigCmd(test.command)
		if !reflect.DeepEqual(got, test.want) || !errors.Is(gotErr, test.wantErr) {
			t.Errorf("ConfigCmd(%v) = %v, %v; want %v, %v", test.command, got, gotErr, test.want, test.wantErr)
		}
	}
}

func TestConfigSet(t *testing.T) {
	tests := []struct {
		command common.Cmd
		want    []byte
		wantErr error
	}{
		// Valid SET
		{common.Cmd{Name: "CONFIG", Args: []string{"SET", "hz", "10"}}, []byte("+OK\r\n"), nil},
		{common.Cmd{Name: "CONFIG", Args: []string{"SET", "active-expire-enabled", "no"}}, []byte("+OK\r\n"), nil},
		{common.Cmd{Name: "CONFIG", Args: []string{"SET", "maxmemory", "1073741824"}}, []byte("+OK\r\n"), nil},
		// Invalid value
		{common.Cmd{Name: "CONFIG", Args: []string{"SET", "hz", "abc"}}, common.RespError("invalid value for config parameter 'hz'"), ErrInvalidConfigValue},
		{common.Cmd{Name: "CONFIG", Args: []string{"SET", "hz", "0"}}, common.RespError("invalid value for config parameter 'hz'"), ErrInvalidConfigValue},
		{common.Cmd{Name: "CONFIG", Args: []string{"SET", "active-expire-enabled", "maybe"}}, common.RespError("invalid value for config parameter 'active-expire-enabled'"), ErrInvalidConfigValue},
		// Unknown parameter
		{common.Cmd{Name: "CONFIG", Args: []string{"SET", "unknown", "value"}}, common.RespError("invalid value for config parameter 'unknown'"), ErrInvalidConfigValue},
		// Wrong number of arguments
		{common.Cmd{Name: "CONFIG", Args: []string{"SET", "hz"}}, common.RespError("wrong number of arguments for 'config' command"), common.ErrWrongNumberOfArgs},
		// Case-insensitive subcommand
		{common.Cmd{Name: "CONFIG", Args: []string{"set", "hz", "10"}}, []byte("+OK\r\n"), nil},
	}
	for _, test := range tests {
		got, gotErr := ConfigCmd(test.command)
		if !reflect.DeepEqual(got, test.want) || !errors.Is(gotErr, test.wantErr) {
			t.Errorf("ConfigCmd(%v) = %v, %v; want %v, %v", test.command, got, gotErr, test.want, test.wantErr)
		}
	}
}

func TestConfigUnknownSubcommand(t *testing.T) {
	tests := []struct {
		command common.Cmd
		want    []byte
	}{
		// Uppercase unknown subcommand
		{common.Cmd{Name: "CONFIG", Args: []string{"REWRITE"}}, common.RespError("unknown subcommand 'REWRITE' for 'config' command")},
		// Lowercase unknown subcommand — uppercased before error message
		{common.Cmd{Name: "CONFIG", Args: []string{"rewrite"}}, common.RespError("unknown subcommand 'REWRITE' for 'config' command")},
	}
	for _, test := range tests {
		got, gotErr := ConfigCmd(test.command)
		if !reflect.DeepEqual(got, test.want) || !errors.Is(gotErr, ErrUnknownSubcommand) {
			t.Errorf("ConfigCmd(%v) = %v, %v; want %v, ErrUnknownSubcommand", test.command, got, gotErr, test.want)
		}
	}
}

func TestConfigSetThenGet(t *testing.T) {
	// Set hz to 5 and verify GET reflects it
	ConfigCmd(common.Cmd{Name: "CONFIG", Args: []string{"SET", "hz", "5"}})
	got, _ := ConfigCmd(common.Cmd{Name: "CONFIG", Args: []string{"GET", "hz"}})
	want := common.RespBulkStringArray([]string{"hz", "5"})
	if !reflect.DeepEqual(got, want) {
		t.Errorf("CONFIG SET then GET: got %v, want %v", got, want)
	}
}
