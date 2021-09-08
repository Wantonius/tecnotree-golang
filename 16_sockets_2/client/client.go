package main 

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	
	conn, err := net.Dial("tcp","localhost:5000")
	if err != nil {
		fmt.Println("Error connecting to server. Reason:",err.Error())
		os.Exit(1)
	}
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("Send message to the server:")
		input,err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Failed to read stdin")
			continue
		}
		conn.Write([]byte(input))
		fmt.Printf("Message sent:%s\n",input)
		//On windows the line end is "\r\n". On others its "\n"
		//Use runtime.GOOS == "windows"
		if strings.TrimRight(input,"\r\n") == "quit" {
			return
		}
	}
}
