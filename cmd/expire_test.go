package cmd

import (
	"errors"
	"reflect"
	"testing"
)

func TestExpire(t *testing.T) {
	// Prepare: insert a key with no TTL
	_, _ = Set([]string{"SET", "tenant", "ACME"})

	tests := []struct {
		command []string
		want    []byte
		wantErr error
	}{
		// Key exists — TTL set, returns 1
		{[]string{"EXPIRE", "tenant", "60"}, []byte(":1\r\n"), nil},
		// Key does not exist — returns 0
		{[]string{"EXPIRE", "ghost", "60"}, []byte(":0\r\n"), nil},
		// Wrong number of arguments
		{[]string{"EXPIRE", "tenant"}, []byte("-wrong number of arguments\r\n"), WrongNumberOfArgumentsError},
		{[]string{"EXPIRE"}, []byte("-wrong number of arguments\r\n"), WrongNumberOfArgumentsError},
		// Non-numeric seconds
		{[]string{"EXPIRE", "tenant", "abc"}, []byte("-invalid expire time in 'expire' command\r\n"), InvalidExpireTimeForExpireError},
		// Zero seconds — must be positive
		{[]string{"EXPIRE", "tenant", "0"}, []byte("-invalid expire time in 'expire' command\r\n"), InvalidExpireTimeForExpireError},
		// Negative seconds
		{[]string{"EXPIRE", "tenant", "-10"}, []byte("-invalid expire time in 'expire' command\r\n"), InvalidExpireTimeForExpireError},
	}
	for _, test := range tests {
		if got, gotErr := Expire(test.command); !reflect.DeepEqual(got, test.want) || !errors.Is(gotErr, test.wantErr) {
			t.Errorf("Expire(%v) = %v, %v; want %v, %v", test.command, got, gotErr, test.want, test.wantErr)
		}
	}
}
