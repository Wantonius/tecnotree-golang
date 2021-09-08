package main

import (
	"net"
	"fmt"
	"bufio"
	"os"
)

func main() {

	connection, _ := net.Dial("tcp","localhost:5000")
	
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Send a message to the server")
	text, _ := reader.ReadString('\n')
	fmt.Fprintf(connection,text+"\n")
	message, _ := bufio.NewReader(connection).ReadString('\n')
	fmt.Printf("Message from the server:%s\n",message)
}