package resp

import (
	"io"
	"net"
	"testing"
)

func TestReplyCommand_WritesBytes(t *testing.T) {
	server, client := net.Pipe()
	defer server.Close()

	ch := make(chan []byte, 1)
	go func() {
		b, _ := io.ReadAll(client)
		ch <- b
	}()

	reply := []byte("+PONG\r\n")
	ReplyCommand(reply, server)
	server.Close()

	got := <-ch
	if string(got) != string(reply) {
		t.Errorf("got %q, want %q", got, reply)
	}
	client.Close()
}

func TestReplyCommand_ClosedConn(t *testing.T) {
	server, client := net.Pipe()
	client.Close()

	// Should not panic even when the connection is closed
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("ReplyCommand panicked on closed connection: %v", r)
		}
	}()
	ReplyCommand([]byte("+PONG\r\n"), server)
	server.Close()
}
