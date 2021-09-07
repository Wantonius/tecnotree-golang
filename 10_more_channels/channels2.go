package main

import (
	"fmt"
	"time"
)

func worker(ch chan string, s time.Duration) {
	time.Sleep(s* time.Millisecond)
	ch <- "worker done"
}

func main() {
	channel := make(chan string)
	channel2 := make(chan string)
	go worker(channel,6500)
	go worker(channel2,3500)
	
	//Select waits for multiple channels and works like a switch-case
	//Default case is for situations where none of the sources have
	//input available. Also used for non-blocking operations.
	L:
	for {
		time.Sleep(time.Second)
		select {
			case v := <-channel: 
				fmt.Println("Worker 1 says",v)
				break L
			case v := <-channel2:
				fmt.Println("Worker 2 says",v)
			default:
				fmt.Println("No input yet")
		}
	}
	
	fmt.Println("---Closing Channels---")
	
	jobs := make(chan int,5)
	done := make(chan bool)
	
	go func() {
		for {
			fmt.Println("Worker: Waiting for more jobs!")
			//Channels have an additional return value which returns
			//true as long as the channel is open and false when
			//the channel is closed
			j,more := <-jobs
			if more {
				fmt.Println("Worker: received job",j)
			} else {
				fmt.Println("Worker: received all jobs")
				done <- true
				return
			}
		}
	}()
	
	for j := 1; j <=3; j++ {
		fmt.Println("Main: Sending another job")
		jobs <- j
		fmt.Println("Main: Sent job")
		time.Sleep(time.Second)
	}
	close(jobs)
	fmt.Println("Main:Sent all jobs and closed channel")
	
	<-done

	fmt.Println("--- Ranging over buffered channels ---")
	
	queue := make(chan string,2)
	queue <- "one"
	queue <- "two"
	close(queue)
	
	fmt.Println("Main: Closed the buffered queue channel")
	
	//Messages in the closed buffered channel will be delivered
	
	for elem := range queue {
		fmt.Printf("Received element %s from a closed channel\n",elem)
	}
}
