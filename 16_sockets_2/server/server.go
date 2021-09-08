package main

import (
	"fmt"
	"net"
	"bufio"
	"os"
)

func handleClient(conn net.Conn) {
	var stringbuffer string
	for {
		buffer,err := bufio.NewReader(conn).ReadBytes('\n')
		if err != nil {
			fmt.Println("Client left")
			conn.Close()
			return
		}
		stringbuffer = string(buffer[:len(buffer)-1])
		fmt.Println("Client Message:",stringbuffer)
	}
}

func main() {
	fmt.Println("Accepting connections at 5000")
	l, err := net.Listen("tcp","localhost:5000")
	if err != nil {
		fmt.Println("Error creating socket and listening for connections. Reason:",err.Error())
		os.Exit(1)
	}
	defer l.Close()
	
	for {
		c, err := l.Accept()
		if err != nil {
			fmt.Println("Error connecting:",err.Error())
			continue;
		}
		fmt.Println("Client "+c.RemoteAddr().String()+" connected")
		go handleClient(c)
	}
}