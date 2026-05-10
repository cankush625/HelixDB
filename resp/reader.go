package resp

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"strconv"
	"strings"
)

// readMessage reads exactly one complete RESP message from the reader and
// returns the raw bytes. It follows the RESP framing so it works correctly
// regardless of how many TCP packets the message arrives in.
//
// Only Array messages are expected for incoming commands. Any other first byte
// is read as a single line and returned as-is so ParseCommand can reject it.
func readMessage(reader *bufio.Reader) ([]byte, error) {
	// Read the first line, e.g. "*3\r\n"
	line, err := reader.ReadString('\n')
	if err != nil {
		return nil, err
	}

	var buf bytes.Buffer
	buf.WriteString(line)

	if len(line) == 0 || line[0] != '*' {
		// Not an array — return the single line for ParseCommand to reject.
		return buf.Bytes(), nil
	}

	// Parse element count from "*N\r\n"
	count, err := strconv.Atoi(strings.TrimRight(line[1:], "\r\n"))
	if err != nil {
		return nil, fmt.Errorf("invalid array length: %w", err)
	}

	for i := 0; i < count; i++ {
		// Read the bulk string length line, e.g. "$5\r\n"
		lenLine, err := reader.ReadString('\n')
		if err != nil {
			return nil, err
		}
		buf.WriteString(lenLine)

		if len(lenLine) == 0 || lenLine[0] != '$' {
			return nil, fmt.Errorf("expected bulk string, got %q", lenLine)
		}

		n, err := strconv.Atoi(strings.TrimRight(lenLine[1:], "\r\n"))
		if err != nil {
			return nil, fmt.Errorf("invalid bulk string length: %w", err)
		}

		// Read exactly n bytes + \r\n
		data := make([]byte, n+2)
		if _, err := io.ReadFull(reader, data); err != nil {
			return nil, err
		}
		buf.Write(data)
	}

	return buf.Bytes(), nil
}
