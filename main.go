package main

import (
	"HelixDB/db"
	"HelixDB/resp"
	"flag"
	"fmt"
	"net"
	"os"
	"os/signal"
	"sync"
	"syscall"
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

	stop := make(chan struct{})
	db.StartActiveExpiry(stop)

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-quit
		fmt.Println("Shutting down...")
		close(stop)
		l.Close()
	}()

	var wg sync.WaitGroup
	for {
		conn, err := l.Accept()
		if err != nil {
			// l.Close() during shutdown causes Accept to return an error.
			// Check if we're shutting down to avoid a spurious log line.
			select {
			case <-stop:
			default:
				fmt.Println("Error accepting connection: ", err.Error())
			}
			break
		}
		wg.Add(1)
		go resp.HandleConn(conn, &wg)
	}

	wg.Wait()
	fmt.Println("Shutdown complete.")
}
