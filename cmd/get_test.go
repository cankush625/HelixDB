package cmd

import (
	"HelixDB/common"
	"HelixDB/db"
	"errors"
	"reflect"
	"testing"
	"time"
)

// TestGet tests the Get function for
// all possible valid and invalid inputs
func TestGet(t *testing.T) {
	// Prepare: insert a persistent key
	_, _ = Set(common.Cmd{Name: "SET", Args: []string{"tenant", "ACME"}})
	// Prepare: insert a key with a far-future TTL (100 seconds)
	_, _ = Set(common.Cmd{Name: "SET", Args: []string{"active_key", "value1", "EX", "100"}})
	// Prepare: insert a key whose TTL is already in the past (expired).
	// Set directly to guarantee the expiration timestamp is behind current time.
	db.DB.Store("expired_key", "value2")
	db.KeyTTL.Store("expired_key", time.Now().UnixMilli()-1000)

	tests := []struct {
		command common.Cmd
		want    []byte
		wantErr error
	}{
		// Key exists in the cache (persistent)
		{common.Cmd{Name: "GET", Args: []string{"tenant"}}, []byte("$4\r\nACME\r\n"), nil},
		// Key doesn't exist in the cache
		{common.Cmd{Name: "GET", Args: []string{"org"}}, []byte("$-1\r\n"), nil},
		// Wrong number of arguments to the GET command
		{common.Cmd{Name: "GET"}, []byte("-wrong number of arguments\r\n"), WrongNumberOfArgumentsError},
		{common.Cmd{Name: "GET", Args: []string{"tenant", "random_arg"}}, []byte("-wrong number of arguments\r\n"), WrongNumberOfArgumentsError},
		// Key exists with a future TTL — should return value
		{common.Cmd{Name: "GET", Args: []string{"active_key"}}, []byte("$6\r\nvalue1\r\n"), nil},
		// Key exists but TTL already expired — should return nil
		{common.Cmd{Name: "GET", Args: []string{"expired_key"}}, []byte("$-1\r\n"), nil},
	}
	for _, test := range tests {
		if got, gotErr := Get(test.command); !reflect.DeepEqual(got, test.want) || !errors.Is(gotErr, test.wantErr) {
			t.Errorf("Get(%v) = %v, %v; want %v, %v", test.command, got, gotErr, test.want, test.wantErr)
		}
	}
}
