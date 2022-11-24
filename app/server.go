package main

import (
	"fmt"
	"net"
	"os"
	// Uncomment this block to pass the first stage
)

func main() {
	// You can use print statements as follows for debugging, they'll be visible when running tests.
	fmt.Println("Logs from your program will appear here!")

	// Uncomment this block to pass the first stage

	l, err := net.Listen("tcp", "0.0.0.0:6379")
	if err != nil {
		fmt.Println("Failed to bind to port 6379")
		os.Exit(1)
	}
	conn, err := l.Accept()
	defer conn.Close()
	if err != nil {
		fmt.Println("Error accepting connection: ", err.Error())
		os.Exit(1)
	}
	go connector(&conn)
}

func connector(c *net.Conn) {
	in := make([]byte, 32768)
	conn := *c
	for {
		if _, err := conn.Read(in); err != nil {
			fmt.Println("Failed to read", err.Error())
		}

		conn.Write([]byte("+PONG\r\n"))
	}
}
