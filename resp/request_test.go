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

func TestPerformRequest_ValidPing(t *testing.T) {
	server, client := newPipe(t)
	ch := readReply(client)

	PerformRequest([]byte("*1\r\n$4\r\nPING\r\n"), server)
	server.Close()

	got := string(<-ch)
	want := "+PONG\r\n"
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func TestPerformRequest_ValidEcho(t *testing.T) {
	server, client := newPipe(t)
	ch := readReply(client)

	PerformRequest([]byte("*2\r\n$4\r\nECHO\r\n$5\r\nhello\r\n"), server)
	server.Close()

	got := string(<-ch)
	want := "$5\r\nhello\r\n"
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func TestPerformRequest_InvalidRESP(t *testing.T) {
	server, client := newPipe(t)
	ch := readReply(client)

	// Unsupported RESP data type prefix
	PerformRequest([]byte("%1\r\n$4\r\nPING\r\n"), server)
	server.Close()

	got := string(<-ch)
	want := "-invalid request\r\n"
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func TestPerformRequest_UnknownCommand(t *testing.T) {
	server, client := newPipe(t)
	ch := readReply(client)

	PerformRequest([]byte("*1\r\n$3\r\nFOO\r\n"), server)
	server.Close()

	got := string(<-ch)
	want := "-unsupported command\r\n"
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func TestPerformRequest_EmptyBuffer(t *testing.T) {
	server, client := newPipe(t)
	ch := readReply(client)

	// nil buffer — ParseCommand returns zero-value Cmd with empty Name, nothing written
	PerformRequest(nil, server)
	server.Close()

	got := string(<-ch)
	if got != "" {
		t.Errorf("expected no reply for empty buffer, got %q", got)
	}
}
