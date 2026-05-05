package main

import (
	"HelixDB/db"
	"HelixDB/resp"
	"flag"
	"fmt"
	"net"
	"os"
)

func main() {
	port := flag.Int("port", 6378, "Port to listen on")
	flag.IntVar(port, "p", *port, "Port to listen on (shorthand)")
	flag.Parse()

	db.ServerConfig.Port = *port

	addr := fmt.Sprintf("0.0.0.0:%d", *port)
	l, err := net.Listen("tcp", addr)
	if err != nil {
		fmt.Printf("Failed to bind to port %d\n", *port)
		os.Exit(1)
	}
	defer l.Close()
	db.StartActiveExpiry()
	for {
		conn, err := l.Accept()
		if err != nil {
			fmt.Println("Error accepting connection: ", err.Error())
			continue
		}
		go resp.HandleConn(conn)
	}
}
