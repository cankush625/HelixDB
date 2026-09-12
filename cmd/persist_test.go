package cmd

import (
	"HelixDB/common"
	"HelixDB/db"
	"errors"
	"reflect"
	"testing"
)

func TestPersist(t *testing.T) {
	// Key with a TTL (~100 seconds from now)
	_, _ = Set(common.Cmd{Name: "SET", Args: []string{"persist_expiring", "v", "EX", "100"}})
	// Key with no TTL
	_, _ = Set(common.Cmd{Name: "SET", Args: []string{"persist_no_ttl", "v"}})
	// Expired key: TTL set to the past but not yet swept by active expiry
	db.DB.Store("persist_expired", db.NewValue("v"))
	db.KeyTTL.Store("persist_expired", int64(1)) // epoch ms well in the past

	tests := []struct {
		command common.Cmd
		want    []byte
		wantErr error
	}{
		// Key with a TTL: the timeout is removed, returns 1
		{common.Cmd{Name: "PERSIST", Args: []string{"persist_expiring"}}, common.RespInteger(1), nil},
		// Key with no TTL: nothing to remove, returns 0
		{common.Cmd{Name: "PERSIST", Args: []string{"persist_no_ttl"}}, common.RespInteger(0), nil},
		// Non-existent key returns 0
		{common.Cmd{Name: "PERSIST", Args: []string{"no_such_key"}}, common.RespInteger(0), nil},
		// Expired-but-unswept key is treated as gone, returns 0
		{common.Cmd{Name: "PERSIST", Args: []string{"persist_expired"}}, common.RespInteger(0), nil},
		// Wrong number of arguments
		{common.Cmd{Name: "PERSIST"}, common.RespError("wrong number of arguments for 'persist' command"), common.ErrWrongNumberOfArgs},
		{common.Cmd{Name: "PERSIST", Args: []string{"a", "b"}}, common.RespError("wrong number of arguments for 'persist' command"), common.ErrWrongNumberOfArgs},
	}
	for _, test := range tests {
		if got, gotErr := Persist(test.command); !reflect.DeepEqual(got, test.want) || !errors.Is(gotErr, test.wantErr) {
			t.Errorf("Persist(%v) = %v, %v; want %v, %v", test.command, got, gotErr, test.want, test.wantErr)
		}
	}

	// The TTL must actually be gone: a second PERSIST on the same key returns 0.
	if got, _ := Persist(common.Cmd{Name: "PERSIST", Args: []string{"persist_expiring"}}); !reflect.DeepEqual(got, common.RespInteger(0)) {
		t.Errorf("Persist(persist_expiring) after removing TTL = %v; want %v (TTL should already be gone)", got, common.RespInteger(0))
	}
}
