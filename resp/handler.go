package resp

import (
	"bufio"
	"io"
	"net"
	"sync"
)

// HandleConn reads RESP commands from the connection and writes replies.
// It uses a bufio.Reader to read one complete RESP message at a time,
// so commands split across multiple TCP packets are handled correctly.
// wg must be held by the caller before invoking HandleConn; it is released
// when the connection is fully done so the server can wait for all
// in-flight connections during graceful shutdown.
func HandleConn(conn net.Conn, wg *sync.WaitGroup) {
	defer wg.Done()
	defer conn.Close()
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
