package main

import (
	"fmt"
	"net"
	"bufio"
	"os"
) 

func main() {

	fmt.Println("Start server at port 5000")
	
	ln,_ := net.Listen("tcp","localhost:5000")
	fmt.Println("Waiting for connection!")
	conn,_ := ln.Accept()
	
	message, _ := bufio.NewReader(conn).ReadString('\n')
	fmt.Println("Message from the client:", string(message))
	fmt.Println("Send a message to the client")
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	fmt.Fprintf(conn,text+"\n")
}