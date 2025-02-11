package resp

import (
	"net"
)

// HandleConn does a read and write to the connection
func HandleConn(conn net.Conn) {
	buffer := make([]byte, 1024)
	for {
		n, err := conn.Read(buffer)
		if err != nil {
			conn.Write([]byte("-Error processing\r\n"))
			return
		}
		if n == 0 {
			conn.Write([]byte("-Error processing\r\n"))
			return
		}
		PerformRequest(buffer[:n], conn)
	}
}
