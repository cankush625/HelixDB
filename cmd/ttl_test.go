package cmd

import (
	"HelixDB/common"
	"HelixDB/db"
	"errors"
	"reflect"
	"testing"
)

func TestTTL(t *testing.T) {
	// Key with no TTL
	_, _ = Set(common.Cmd{Name: "SET", Args: []string{"ttl_persist", "v"}})
	// Key with TTL (~100 seconds from now)
	_, _ = Set(common.Cmd{Name: "SET", Args: []string{"ttl_expiring", "v", "EX", "100"}})
	// Expired key (TTL set to the past)
	db.DB.Store("ttl_expired", db.NewValue("v"))
	db.KeyTTL.Store("ttl_expired", int64(1)) // epoch ms well in the past

	tests := []struct {
		command common.Cmd
		want    []byte
		wantErr error
	}{
		// Key with no TTL returns -1
		{common.Cmd{Name: "TTL", Args: []string{"ttl_persist"}}, common.RespInteger(-1), nil},
		// Non-existent key returns -2
		{common.Cmd{Name: "TTL", Args: []string{"no_such_key"}}, common.RespInteger(-2), nil},
		// Expired key returns -2
		{common.Cmd{Name: "TTL", Args: []string{"ttl_expired"}}, common.RespInteger(-2), nil},
		// Wrong number of arguments
		{common.Cmd{Name: "TTL"}, common.RespError("wrong number of arguments for 'ttl' command"), common.ErrWrongNumberOfArgs},
		{common.Cmd{Name: "TTL", Args: []string{"a", "b"}}, common.RespError("wrong number of arguments for 'ttl' command"), common.ErrWrongNumberOfArgs},
	}
	for _, test := range tests {
		if got, gotErr := TTL(test.command); !reflect.DeepEqual(got, test.want) || !errors.Is(gotErr, test.wantErr) {
			t.Errorf("TTL(%v) = %v, %v; want %v, %v", test.command, got, gotErr, test.want, test.wantErr)
		}
	}

	// Key with TTL returns a positive integer (exact value is timing-sensitive)
	got, gotErr := TTL(common.Cmd{Name: "TTL", Args: []string{"ttl_expiring"}})
	if gotErr != nil {
		t.Errorf("TTL(ttl_expiring) unexpected error: %v", gotErr)
	}
	if reflect.DeepEqual(got, common.RespInteger(-1)) || reflect.DeepEqual(got, common.RespInteger(-2)) {
		t.Errorf("TTL(ttl_expiring) = %v; want a positive integer", got)
	}
}

func TestPTTL(t *testing.T) {
	// Key with no TTL
	_, _ = Set(common.Cmd{Name: "SET", Args: []string{"pttl_persist", "v"}})
	// Key with TTL (~100 seconds from now)
	_, _ = Set(common.Cmd{Name: "SET", Args: []string{"pttl_expiring", "v", "EX", "100"}})
	// Expired key
	db.DB.Store("pttl_expired", db.NewValue("v"))
	db.KeyTTL.Store("pttl_expired", int64(1))

	tests := []struct {
		command common.Cmd
		want    []byte
		wantErr error
	}{
		// Key with no TTL returns -1
		{common.Cmd{Name: "PTTL", Args: []string{"pttl_persist"}}, common.RespInteger(-1), nil},
		// Non-existent key returns -2
		{common.Cmd{Name: "PTTL", Args: []string{"no_such_key"}}, common.RespInteger(-2), nil},
		// Expired key returns -2
		{common.Cmd{Name: "PTTL", Args: []string{"pttl_expired"}}, common.RespInteger(-2), nil},
		// Wrong number of arguments
		{common.Cmd{Name: "PTTL"}, common.RespError("wrong number of arguments for 'pttl' command"), common.ErrWrongNumberOfArgs},
		{common.Cmd{Name: "PTTL", Args: []string{"a", "b"}}, common.RespError("wrong number of arguments for 'pttl' command"), common.ErrWrongNumberOfArgs},
	}
	for _, test := range tests {
		if got, gotErr := PTTL(test.command); !reflect.DeepEqual(got, test.want) || !errors.Is(gotErr, test.wantErr) {
			t.Errorf("PTTL(%v) = %v, %v; want %v, %v", test.command, got, gotErr, test.want, test.wantErr)
		}
	}

	// Key with TTL returns a positive integer
	got, gotErr := PTTL(common.Cmd{Name: "PTTL", Args: []string{"pttl_expiring"}})
	if gotErr != nil {
		t.Errorf("PTTL(pttl_expiring) unexpected error: %v", gotErr)
	}
	if reflect.DeepEqual(got, common.RespInteger(-1)) || reflect.DeepEqual(got, common.RespInteger(-2)) {
		t.Errorf("PTTL(pttl_expiring) = %v; want a positive integer", got)
	}
}
