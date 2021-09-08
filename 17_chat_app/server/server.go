package main

import (
	"fmt"
	"net"
	"bufio"
	"strings"
	"os"
)

type Manager struct {
	broadcast	chan string
	register	chan *Client
	unregister	chan *Client
	clients		map[*Client]bool
}

type Client struct {
	socket 	net.Conn
	data 	chan string
	name	string
}

func (manager *Manager) startService() {
	for {
		select {
			case connection := <-manager.register:
				manager.clients[connection] = true
				fmt.Printf("User %s has entered the chat\n",connection.name)
				for conn := range manager.clients {
					select {
						case conn.data <- "User "+connection.name+" has entered\n":
						default: 
							close(conn.data)
							delete(manager.clients,conn)
						}
				}		
			case connection := <- manager.unregister:
				if _,ok := manager.clients[connection]; ok {
					close(connection.data)
					fmt.Printf("User %s has left the chat\n",connection.name)
					delete(manager.clients,connection)
					for conn := range manager.clients {
						select {
							case conn.data <- "User "+connection.name+" has left\n":
							default: 
								close(conn.data)
								delete(manager.clients,conn)
							}
					}				
				}
			case message := <-manager.broadcast:
				for conn := range manager.clients{
					select {
						case conn.data <- message:
						default: 
							close(conn.data)
							delete(manager.clients,conn)
						}			
				}
		}	
	}
}

func (manager *Manager) receive(client *Client) {
	for {
		message := make([]byte,256)
		length,err := client.socket.Read(message)
		if err != nil {
			manager.unregister <- client
			client.socket.Close()
			return
		}
		if length > 0 {
			temp_message := client.name+":"+string(message)
			manager.broadcast <- temp_message
		}
	}

}

func (manager *Manager) send(client *Client) {
	defer client.socket.Close()
	for {
		select {
			case message, ok := <-client.data:
				if !ok {
					return
				}
				fmt.Fprintf(client.socket,message)
		}
	}
}

func main() {
	fmt.Println("Starting chat server at 5000")
	listener, err := net.Listen("tcp","localhost:5000")
	if err != nil {
		fmt.Printf("Failure to bind and listen. Reason:%s\n",err.Error())
		os.Exit(1)
	}
	manager := Manager{
		broadcast:		make(chan string),
		register:		make(chan *Client),
		unregister:		make(chan *Client),
		clients:		make(map[*Client]bool),
	}
	go manager.startService()
	for {
		fmt.Println("Waiting for a new client")
		connection,err := listener.Accept()
		if err != nil {
			fmt.Println("Failed to accept connection. Reason:",err.Error())
			continue
		}
		fmt.Println("New client! Waiting for a name")
		client_name, err := bufio.NewReader(connection).ReadString('\n')
		if err != nil {
			fmt.Println("Failed to read name. Reason:",err.Error())
			connection.Close()
			continue		
		}
		client_name = strings.TrimRight(client_name,"\r\n")
		client := &Client{socket:connection,data:make(chan string),name:client_name}
		manager.register <- client
		go manager.receive(client)
		go manager.send(client)
	}

}
 
