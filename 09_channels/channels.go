package main

import (
	"fmt"
	"time"
)

func worker(done chan bool) {
	fmt.Println("Worker: Lets do some work!")
	time.Sleep(3*time.Second)
	fmt.Println("Worker: And we are done")
	fmt.Println("Worker: Sending the done signal to main")
	done <- true
}

func main() {

	messages := make(chan string)
	
	//Create a channel with make(chan val-type). Channels are typed
	//by the value. Sends and recieves block until both are ready

	fmt.Println("---Basic channel---")
	
	go func() {
		fmt.Println("Pinger:Pinging the main")
		messages <- "ping"
	}()
	
	fmt.Println("Waiting for ping!")
	msg := <- messages
	fmt.Printf("%s\n",msg)
	
	time.Sleep(2*time.Second)
	
	fmt.Println("---Buffered channel---")
	
	buffered := make(chan string,2)
	
	buffered <- "buffered"
	buffered <- "channel"
	
	fmt.Println(<-buffered)
	fmt.Println(<-buffered)
	time.Sleep(2*time.Second)
	
	//Lets use channels to synchro execution across goroutines
	done := make(chan bool,1)
	go worker(done)
	fmt.Println("Main: Waiting for the worker to complete")
	<-done
	fmt.Println("Main: Worker done. Exiting!")
}
