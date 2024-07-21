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
		buffer = buffer[:n]
		fmt.Printf("buffer: %s", buffer)
		if n == 0 {
			return
		}
		_, err = conn.Write([]byte("+PONG\r\n"))
		if err != nil {
			err := fmt.Errorf("error writing to a connection")
			fmt.Println(err.Error())
		}
	}
}
