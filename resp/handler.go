package resp

import (
	"fmt"
	"net"
)

// HandleConn does a read and write to the connection
func HandleConn(conn net.Conn) {
	buffer := make([]byte, 1024)
	for {
		n, err := conn.Read(buffer)
		if err != nil {
			return
		}
		buffer = buffer[:n]
		fmt.Printf("buffer: %s", buffer)
		if n == 0 {
			return
		}
		PerformRequest(buffer, conn)
	}
}
