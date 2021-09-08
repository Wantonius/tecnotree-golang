package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"
	"os"
) 

func send(conn net.Conn, done chan bool, reader bufio.Reader) {
	for {
		fmt.Print("Send message to chat:")
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Can't read from stdin")
			os.Exit(1)
		}
		conn.Write([]byte(input))
		fmt.Printf("Message sent:%s\n",input)
		if strings.TrimRight(input,"\r\n") == "quit" {
			done <- true
			conn.Close()
			os.Exit(1)
		}	
	}
}

func receive(conn net.Conn,done chan bool) {
	var stringbuffer string
	for {
		buffer,err := bufio.NewReader(conn).ReadBytes('\n')
		if err != nil {
			fmt.Println("Problem with the server")
			conn.Close()
			done <- true
			os.Exit(1)
		}
		stringbuffer = string(buffer[:len(buffer)-1])
		fmt.Printf("\nServer message:%s\n",stringbuffer)
	}

}

func main() {
	var done chan bool
	fmt.Println("Welcome to chat app")
	fmt.Println("Please enter your name")
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Can't read from stdin")
		os.Exit(1)
	}
	conn,err := net.Dial("tcp","localhost:5000")
	if err != nil {
		fmt.Println("Error connecting to server. Reason:",err.Error())
		os.Exit(1)
	}
	conn.Write([]byte(input))
	go send(conn,done,*reader)
	go receive(conn,done)
	<-done

}