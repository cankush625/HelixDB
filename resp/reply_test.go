package resp

import (
	"io"
	"net"
	"testing"
)

// TestReplyCommand tests the ReplyCommand function for
// all possible valid and invalid inputs
func TestReplyCommand(t *testing.T) {
	tests := []struct {
		reply []byte
		want  string
	}{
		{[]byte("+PONG\r\n"), "+PONG\r\n"},
		{[]byte("+OK\r\n"), "+OK\r\n"},
		{[]byte("$5\r\nhello\r\n"), "$5\r\nhello\r\n"},
		{[]byte("-error message\r\n"), "-error message\r\n"},
	}
	for _, test := range tests {
		server, client := net.Pipe()

		ch := make(chan []byte, 1)
		go func() {
			b, _ := io.ReadAll(client)
			ch <- b
		}()

		ReplyCommand(test.reply, server)
		server.Close()

		got := string(<-ch)
		if got != test.want {
			t.Errorf("ReplyCommand(%q) = %q; want %q", test.reply, got, test.want)
		}
		client.Close()
	}
}

// TestReplyCommand_ClosedConn tests that ReplyCommand does not panic
// when the connection is already closed
func TestReplyCommand_ClosedConn(t *testing.T) {
	server, client := net.Pipe()
	client.Close()

	defer func() {
		if r := recover(); r != nil {
			t.Errorf("ReplyCommand panicked on closed connection: %v", r)
		}
	}()
	ReplyCommand([]byte("+PONG\r\n"), server)
	server.Close()
}
