package resp

import (
	"bufio"
	"io"
	"net"
)

// HandleConn reads RESP commands from the connection and writes replies.
// It uses a bufio.Reader to read one complete RESP message at a time,
// so commands split across multiple TCP packets are handled correctly.
func HandleConn(conn net.Conn) {
	reader := bufio.NewReader(conn)
	for {
		msg, err := readMessage(reader)
		if err != nil {
			// io.EOF means the client closed the connection cleanly.
			if err == io.EOF {
				return
			}
			conn.Write([]byte("-Error processing\r\n"))
			return
		}
		PerformRequest(msg, conn)
	}
}
