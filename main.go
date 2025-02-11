package main

import (
	"HelixDB/resp"
	"fmt"
	"net"
	"os"
)

func main() {
	l, err := net.Listen("tcp", "0.0.0.0:6378")
	if err != nil {
		fmt.Println("Failed to bind to port 6378")
		os.Exit(1)
	}
	defer l.Close()
	for {
		conn, err := l.Accept()
		if err != nil {
			fmt.Println("Error accepting connection: ", err.Error())
			continue
		}
		go resp.HandleConn(conn)
	}
}
