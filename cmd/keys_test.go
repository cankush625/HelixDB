package cmd

import (
	"HelixDB/common"
	"HelixDB/db"
	"bytes"
	"errors"
	"sort"
	"strings"
	"testing"
	"time"
)

func TestKeys(t *testing.T) {
	// Clear state
	db.DB.Clear()
	db.KeyTTL.Clear()

	// Persistent keys
	_, _ = Set(common.Cmd{Name: "SET", Args: []string{"hello", "1"}})
	_, _ = Set(common.Cmd{Name: "SET", Args: []string{"hallo", "2"}})
	_, _ = Set(common.Cmd{Name: "SET", Args: []string{"hxllo", "3"}})
	_, _ = Set(common.Cmd{Name: "SET", Args: []string{"world", "4"}})

	// Expired key — should not appear in results
	db.DB.Store("expired", "x")
	db.KeyTTL.Store("expired", time.Now().UnixMilli()-1000)

	tests := []struct {
		command      common.Cmd
		wantKeys     []string // expected keys, order-independent
		wantErr      error
	}{
		// Match all keys
		{common.Cmd{Name: "KEYS", Args: []string{"*"}}, []string{"hello", "hallo", "hxllo", "world"}, nil},
		// ? matches single character
		{common.Cmd{Name: "KEYS", Args: []string{"h?llo"}}, []string{"hello", "hallo", "hxllo"}, nil},
		// Exact match
		{common.Cmd{Name: "KEYS", Args: []string{"world"}}, []string{"world"}, nil},
		// No match
		{common.Cmd{Name: "KEYS", Args: []string{"nomatch"}}, []string{}, nil},
		// Character class
		{common.Cmd{Name: "KEYS", Args: []string{"h[ae]llo"}}, []string{"hello", "hallo"}, nil},
		// Wrong number of arguments
		{common.Cmd{Name: "KEYS"}, nil, common.ErrWrongNumberOfArgs},
		{common.Cmd{Name: "KEYS", Args: []string{"*", "extra"}}, nil, common.ErrWrongNumberOfArgs},
		// Invalid pattern — unclosed bracket
		{common.Cmd{Name: "KEYS", Args: []string{"["}}, nil, SyntaxError},
	}

	for _, test := range tests {
		got, gotErr := Keys(test.command)
		if !errors.Is(gotErr, test.wantErr) {
			t.Errorf("Keys(%v) error = %v; want %v", test.command, gotErr, test.wantErr)
			continue
		}
		if test.wantErr != nil {
			continue
		}
		if !matchBulkStringArray(got, test.wantKeys) {
			t.Errorf("Keys(%v) = %s; want keys %v", test.command, got, test.wantKeys)
		}
	}
}

// matchBulkStringArray checks that a command response bytes
// contain exactly the expected keys as bulk strings, regardless of order.
func matchBulkStringArray(rawResponse []byte, expectedKeys []string) bool {
	lines := strings.Split(string(rawResponse), "\r\n")
	if len(lines) == 0 {
		return false
	}
	// Extract values from bulk string entries (lines starting with $)
	var gotKeys []string
	for i, line := range lines {
		if len(line) > 0 && line[0] == '$' && i+1 < len(lines) {
			gotKeys = append(gotKeys, lines[i+1])
		}
	}
	if len(gotKeys) != len(expectedKeys) {
		return false
	}
	sort.Strings(gotKeys)
	sort.Strings(expectedKeys)
	return bytes.Equal([]byte(strings.Join(gotKeys, ",")), []byte(strings.Join(expectedKeys, ",")))
}
