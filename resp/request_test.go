package resp

import (
	"io"
	"net"
	"testing"
)

// newPipe returns a connected pair of net.Conn. server is passed to the function
// under test; client is used to read back what was written.
func newPipe(t *testing.T) (server, client net.Conn) {
	t.Helper()
	server, client = net.Pipe()
	t.Cleanup(func() {
		server.Close()
		client.Close()
	})
	return
}

// readReply reads from the client side of the pipe in a goroutine and returns
// the bytes via a channel. The server side must be closed (or the write must
// complete) before this unblocks.
func readReply(client net.Conn) <-chan []byte {
	ch := make(chan []byte, 1)
	go func() {
		b, _ := io.ReadAll(client)
		ch <- b
	}()
	return ch
}

// TestPerformRequest tests the PerformRequest function for
// all possible valid and invalid inputs
func TestPerformRequest(t *testing.T) {
	tests := []struct {
		input []byte
		want  string
	}{
		// Valid commands
		{[]byte("*1\r\n$4\r\nPING\r\n"), "+PONG\r\n"},
		{[]byte("*2\r\n$4\r\nECHO\r\n$5\r\nhello\r\n"), "$5\r\nhello\r\n"},
		// Invalid RESP prefix
		{[]byte("%1\r\n$4\r\nPING\r\n"), "-invalid request\r\n"},
		// Unknown command
		{[]byte("*1\r\n$3\r\nFOO\r\n"), "-unsupported command\r\n"},
		// nil buffer — ParseCommand returns zero-value Cmd with empty Name, nothing written
		{nil, ""},
	}
	for _, test := range tests {
		server, client := newPipe(t)
		ch := readReply(client)

		PerformRequest(test.input, server)
		server.Close()

		got := string(<-ch)
		if got != test.want {
			t.Errorf("PerformRequest(%q) = %q; want %q", test.input, got, test.want)
		}
	}
}
