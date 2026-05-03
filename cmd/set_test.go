package cmd

import (
	"HelixDB/common"
	"HelixDB/db"
	"errors"
	"fmt"
	"reflect"
	"testing"
	"time"
)

// TestSet tests the Set function for all possible valid and invalid inputs.
func TestSet(t *testing.T) {
	tests := []struct {
		command common.Cmd
		want    []byte
		wantErr error
	}{
		// Basic SET
		{common.Cmd{Name: "SET", Args: []string{"tenant", "ACME"}}, []byte("+OK\r\n"), nil},
		// Missing key/value
		{common.Cmd{Name: "SET", Args: []string{"tenant"}}, []byte("-wrong number of arguments for 'set' command\r\n"), common.ErrWrongNumberOfArgs},
		{common.Cmd{Name: "SET"}, []byte("-wrong number of arguments for 'set' command\r\n"), common.ErrWrongNumberOfArgs},
		// Unknown option
		{common.Cmd{Name: "SET", Args: []string{"tenant", "ACME", "asd"}}, []byte("-syntax error\r\n"), SyntaxError},
		// Valid EX
		{common.Cmd{Name: "SET", Args: []string{"tenant", "ACME", "EX", "10"}}, []byte("+OK\r\n"), nil},
		// Valid PX
		{common.Cmd{Name: "SET", Args: []string{"tenant", "ACME", "PX", "5000"}}, []byte("+OK\r\n"), nil},
		// EX and PX together — mutually exclusive
		{common.Cmd{Name: "SET", Args: []string{"tenant", "ACME", "EX", "10", "PX", "5000"}}, []byte("-syntax error\r\n"), SyntaxError},
		// EX with non-numeric value
		{common.Cmd{Name: "SET", Args: []string{"tenant", "ACME", "EX", "abc"}}, []byte("-invalid expire time in 'set' command\r\n"), InvalidExpireTimeError},
		// EX with zero — must be positive
		{common.Cmd{Name: "SET", Args: []string{"tenant", "ACME", "EX", "0"}}, []byte("-invalid expire time in 'set' command\r\n"), InvalidExpireTimeError},
		// EX with negative value
		{common.Cmd{Name: "SET", Args: []string{"tenant", "ACME", "EX", "-5"}}, []byte("-invalid expire time in 'set' command\r\n"), InvalidExpireTimeError},
		// PX with non-numeric value
		{common.Cmd{Name: "SET", Args: []string{"tenant", "ACME", "PX", "abc"}}, []byte("-invalid expire time in 'set' command\r\n"), InvalidExpireTimeError},
		// PX with zero — must be positive
		{common.Cmd{Name: "SET", Args: []string{"tenant", "ACME", "PX", "0"}}, []byte("-invalid expire time in 'set' command\r\n"), InvalidExpireTimeError},
		// EX missing its value
		{common.Cmd{Name: "SET", Args: []string{"tenant", "ACME", "EX"}}, []byte("-syntax error\r\n"), SyntaxError},
		// Valid EXAT
		{common.Cmd{Name: "SET", Args: []string{"tenant", "ACME", "EXAT", "9999999999"}}, []byte("+OK\r\n"), nil},
		// Valid PXAT
		{common.Cmd{Name: "SET", Args: []string{"tenant", "ACME", "PXAT", "9999999999000"}}, []byte("+OK\r\n"), nil},
		// EXAT with invalid value
		{common.Cmd{Name: "SET", Args: []string{"tenant", "ACME", "EXAT", "abc"}}, []byte("-invalid expire time in 'set' command\r\n"), InvalidExpireTimeError},
		// EXAT with zero
		{common.Cmd{Name: "SET", Args: []string{"tenant", "ACME", "EXAT", "0"}}, []byte("-invalid expire time in 'set' command\r\n"), InvalidExpireTimeError},
		// PXAT with invalid value
		{common.Cmd{Name: "SET", Args: []string{"tenant", "ACME", "PXAT", "abc"}}, []byte("-invalid expire time in 'set' command\r\n"), InvalidExpireTimeError},
		// EX and EXAT together — mutually exclusive
		{common.Cmd{Name: "SET", Args: []string{"tenant", "ACME", "EX", "10", "EXAT", "9999999999"}}, []byte("-syntax error\r\n"), SyntaxError},
		// NX and XX together — mutually exclusive
		{common.Cmd{Name: "SET", Args: []string{"tenant", "ACME", "NX", "XX"}}, []byte("-syntax error\r\n"), SyntaxError},
		// EX and KEEPTTL together — mutually exclusive
		{common.Cmd{Name: "SET", Args: []string{"tenant", "ACME", "EX", "10", "KEEPTTL"}}, []byte("-syntax error\r\n"), SyntaxError},
	}
	for _, test := range tests {
		if got, gotErr := Set(test.command); !reflect.DeepEqual(got, test.want) || !errors.Is(gotErr, test.wantErr) {
			t.Errorf("Set(%v) = %v, %v; want %v, %v", test.command, got, gotErr, test.want, test.wantErr)
		}
	}
}

func TestSetNX(t *testing.T) {
	db.DB.Delete("nx_key")
	db.KeyTTL.Delete("nx_key")

	// NX on a non-existent key — should set
	got, err := Set(common.Cmd{Name: "SET", Args: []string{"nx_key", "value1", "NX"}})
	if !reflect.DeepEqual(got, []byte("+OK\r\n")) || err != nil {
		t.Errorf("SET NX on new key: got %v, %v; want +OK, nil", got, err)
	}

	// NX on an existing key — should not set, return nil
	got, err = Set(common.Cmd{Name: "SET", Args: []string{"nx_key", "value2", "NX"}})
	if !reflect.DeepEqual(got, common.RespNullBulkString()) || err != nil {
		t.Errorf("SET NX on existing key: got %v, %v; want nil bulk string, nil", got, err)
	}

	// Value should still be the original
	val, _ := GetValueFromMemory("nx_key")
	if val != "value1" {
		t.Errorf("SET NX overwrote existing key: got %v, want value1", val)
	}
}

func TestSetXX(t *testing.T) {
	db.DB.Delete("xx_key")
	db.KeyTTL.Delete("xx_key")

	// XX on a non-existent key — should not set, return nil
	got, err := Set(common.Cmd{Name: "SET", Args: []string{"xx_key", "value1", "XX"}})
	if !reflect.DeepEqual(got, common.RespNullBulkString()) || err != nil {
		t.Errorf("SET XX on non-existent key: got %v, %v; want nil bulk string, nil", got, err)
	}

	// Seed the key
	Set(common.Cmd{Name: "SET", Args: []string{"xx_key", "original"}})

	// XX on an existing key — should set
	got, err = Set(common.Cmd{Name: "SET", Args: []string{"xx_key", "updated", "XX"}})
	if !reflect.DeepEqual(got, []byte("+OK\r\n")) || err != nil {
		t.Errorf("SET XX on existing key: got %v, %v; want +OK, nil", got, err)
	}

	val, _ := GetValueFromMemory("xx_key")
	if val != "updated" {
		t.Errorf("SET XX did not update value: got %v, want updated", val)
	}
}

func TestSetGET(t *testing.T) {
	db.DB.Delete("get_key")
	db.KeyTTL.Delete("get_key")

	// GET on a non-existent key — should return nil bulk string
	got, err := Set(common.Cmd{Name: "SET", Args: []string{"get_key", "first", "GET"}})
	if !reflect.DeepEqual(got, common.RespNullBulkString()) || err != nil {
		t.Errorf("SET GET on new key: got %v, %v; want nil bulk string, nil", got, err)
	}

	// GET on an existing key — should return old value
	got, err = Set(common.Cmd{Name: "SET", Args: []string{"get_key", "second", "GET"}})
	if !reflect.DeepEqual(got, common.RespBulkString("first")) || err != nil {
		t.Errorf("SET GET on existing key: got %v, %v; want bulk string 'first', nil", got, err)
	}
}

func TestSetKEEPTTL(t *testing.T) {
	db.DB.Delete("ttl_key")
	db.KeyTTL.Delete("ttl_key")

	// Set key with a TTL
	Set(common.Cmd{Name: "SET", Args: []string{"ttl_key", "original", "PX", "60000"}})
	expiry, _ := db.KeyTTL.Load("ttl_key")

	// Overwrite with KEEPTTL — TTL should be preserved
	Set(common.Cmd{Name: "SET", Args: []string{"ttl_key", "updated", "KEEPTTL"}})
	newExpiry, _ := db.KeyTTL.Load("ttl_key")

	if expiry != newExpiry {
		t.Errorf("KEEPTTL changed the TTL: got %v, want %v", newExpiry, expiry)
	}

	// Value should be updated
	val, _ := GetValueFromMemory("ttl_key")
	if val != "updated" {
		t.Errorf("KEEPTTL did not update value: got %v, want updated", val)
	}
}

func TestSetEXAT(t *testing.T) {
	futureTimestamp := time.Now().Add(time.Hour).Unix()
	got, err := Set(common.Cmd{Name: "SET", Args: []string{"exat_key", "value", "EXAT", itoa(futureTimestamp)}})
	if !reflect.DeepEqual(got, []byte("+OK\r\n")) || err != nil {
		t.Errorf("SET EXAT: got %v, %v; want +OK, nil", got, err)
	}

	val, err := GetValueFromMemory("exat_key")
	if err != nil || val != "value" {
		t.Errorf("SET EXAT key not accessible: got %v, %v", val, err)
	}
}

func TestSetPXAT(t *testing.T) {
	futureTimestamp := time.Now().Add(time.Hour).UnixMilli()
	got, err := Set(common.Cmd{Name: "SET", Args: []string{"pxat_key", "value", "PXAT", itoa(futureTimestamp)}})
	if !reflect.DeepEqual(got, []byte("+OK\r\n")) || err != nil {
		t.Errorf("SET PXAT: got %v, %v; want +OK, nil", got, err)
	}

	val, err := GetValueFromMemory("pxat_key")
	if err != nil || val != "value" {
		t.Errorf("SET PXAT key not accessible: got %v, %v", val, err)
	}
}

func itoa(n int64) string {
	return fmt.Sprintf("%d", n)
}
