package main

import (
	"fmt"
	"io"
	"net"
	"os"
	"unicode"

	// Uncomment this block to pass the first stage

	"strings"
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
	for {
		conn, err := l.Accept()
		defer conn.Close()
		if err != nil {
			fmt.Println("Error accepting connection: ", err.Error())
			os.Exit(1)
		}
		go connector(&conn)
	}
}

func connector(c *net.Conn) {
	in := make([]byte, 32768)
	conn := *c
	for {
		if _, err := conn.Read(in); err != nil {
			if err == io.EOF {
				break
			} else {
				fmt.Println("Failed to read", err.Error())
			}
		}

		if in[0] == '*' && unicode.IsNumber(rune(in[1])) {
			stin := string(in)
			cmd := strings.ReplaceAll(stin, "\r\n", " ")
			cmdArr := strings.Split(cmd, " ")
			if cmdArr[2] == "ECHO" || cmdArr[2] == "echo" {
				fmt.Println(cmdArr)
				conn.Write([]byte(fmt.Sprintf("+%s\r\n", cmdArr[4])))
			} else if cmdArr[2] == "ping" || cmdArr[2] == "PING" {
				conn.Write([]byte("+PONG\r\n"))
			}
		}

	}
}
