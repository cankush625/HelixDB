package resp

import (
	"bufio"
	"bytes"
	"io"
	"strings"
	"testing"
)

// TestReadMessage tests the readMessage function for complete and partial reads.
func TestReadMessage(t *testing.T) {
	tests := []struct {
		input   string
		want    string
		wantErr bool
	}{
		// Single-element array (PING)
		{"*1\r\n$4\r\nPING\r\n", "*1\r\n$4\r\nPING\r\n", false},
		// Two-element array (ECHO hello)
		{"*2\r\n$4\r\nECHO\r\n$5\r\nhello\r\n", "*2\r\n$4\r\nECHO\r\n$5\r\nhello\r\n", false},
		// Three-element array (SET key value)
		{"*3\r\n$3\r\nSET\r\n$3\r\nkey\r\n$5\r\nvalue\r\n", "*3\r\n$3\r\nSET\r\n$3\r\nkey\r\n$5\r\nvalue\r\n", false},
		// Non-array first byte — returned as-is for ParseCommand to reject
		{"+OK\r\n", "+OK\r\n", false},
	}
	for _, test := range tests {
		reader := bufio.NewReader(strings.NewReader(test.input))
		got, err := readMessage(reader)
		if test.wantErr && err == nil {
			t.Errorf("readMessage(%q) expected error, got nil", test.input)
			continue
		}
		if !test.wantErr && err != nil {
			t.Errorf("readMessage(%q) unexpected error: %v", test.input, err)
			continue
		}
		if string(got) != test.want {
			t.Errorf("readMessage(%q) = %q; want %q", test.input, got, test.want)
		}
	}
}

// TestReadMessage_MultipleCommands tests that readMessage reads one command at
// a time from a stream containing multiple pipelined commands.
func TestReadMessage_MultipleCommands(t *testing.T) {
	// Two commands sent back-to-back in the same stream (pipelining)
	input := "*1\r\n$4\r\nPING\r\n" +
		"*2\r\n$4\r\nECHO\r\n$5\r\nhello\r\n"

	reader := bufio.NewReader(strings.NewReader(input))

	first, err := readMessage(reader)
	if err != nil {
		t.Fatalf("first readMessage error: %v", err)
	}
	if string(first) != "*1\r\n$4\r\nPING\r\n" {
		t.Errorf("first message = %q; want %q", first, "*1\r\n$4\r\nPING\r\n")
	}

	second, err := readMessage(reader)
	if err != nil {
		t.Fatalf("second readMessage error: %v", err)
	}
	if string(second) != "*2\r\n$4\r\nECHO\r\n$5\r\nhello\r\n" {
		t.Errorf("second message = %q; want %q", second, "*2\r\n$4\r\nECHO\r\n$5\r\nhello\r\n")
	}

	// Third read should return EOF
	_, err = readMessage(reader)
	if err != io.EOF {
		t.Errorf("expected io.EOF after last command, got %v", err)
	}
}

// TestReadMessage_SplitAcrossReads simulates a command arriving in two separate
// TCP chunks by writing to a pipe with a delay between writes.
func TestReadMessage_SplitAcrossReads(t *testing.T) {
	// Simulate split delivery: write two parts with a pipe
	pr, pw := io.Pipe()
	reader := bufio.NewReader(pr)

	done := make(chan []byte, 1)
	go func() {
		msg, _ := readMessage(reader)
		done <- msg
	}()

	// Write the first half
	pw.Write([]byte("*3\r\n$3\r\nSET\r\n"))
	// Write the second half
	pw.Write([]byte("$3\r\nkey\r\n$5\r\nvalue\r\n"))
	pw.Close()

	got := <-done
	want := "*3\r\n$3\r\nSET\r\n$3\r\nkey\r\n$5\r\nvalue\r\n"
	if !bytes.Equal(got, []byte(want)) {
		t.Errorf("split read = %q; want %q", got, want)
	}
}
