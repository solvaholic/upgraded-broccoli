package main

import (
	"io"
	"log"
	"net"
)

// echo is a handler function that simply echoes received data.
func echo(conn net.Conn) {
	// Copy data from io.Reader to io.Writer
	if _, err := io.Copy(conn, conn); err != nil {
		log.Fatalln("Unable to read/write data")
	}
	conn.Close()
	log.Println("Connection closed")
}

func main() {
	// Bind to TCP port 20080 on all interfaces.
	listener, err := net.Listen("tcp", "127.0.0.1:20080")
	if err != nil {
		log.Fatalln("Unable to bind to port")
	}
	log.Println("Listening on 127.0.0.1:20080")
	for {
		// Wait for connection. Create net.Conn on connection established.
		conn, err := listener.Accept()
		log.Println("Received connection")
		if err != nil {
			log.Fatalln("Unable to accept connection")
		}
		// Handle the connection. Using goroutine for concurrency.
		go echo(conn)
	}
}
