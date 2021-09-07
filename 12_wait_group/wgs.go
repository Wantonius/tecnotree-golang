package main

import (
	"fmt"
	"sync"
	"time"
	"math/rand"
)

func worker(id int, wg *sync.WaitGroup) {
	
	defer wg.Done()
	var temp int
	
	rand.Seed(time.Now().UnixNano())
	
	temp = rand.Intn(10)
	temp_time := time.Duration(temp)
	fmt.Printf("Worker %d: Starting job. Takes %d seconds\n",id,temp)
	time.Sleep(temp_time*time.Second)
	fmt.Printf("Worker %d: Ending job\n",id)
}

func control(wg *sync.WaitGroup) {

	var control sync.WaitGroup
	fmt.Println("Control: Setting up workers")
	
	for i := 1;i<=10;i++ {
		control.Add(1)
		go worker(i,&control)
		time.Sleep(time.Second)
	}

	fmt.Println("Control: Waiting for the jobs the be completed")
	
	control.Wait()
	
	fmt.Println("Control: Jobs done. Releasing Main")
	
	wg.Done()
}

func main() {

	var wg sync.WaitGroup
	
	fmt.Println("Main: Starting control")
	wg.Add(1)
	go control(&wg)
	
	fmt.Println("Main: Waiting on control")
	
	wg.Wait()
	
	fmt.Println("Main: Control done. Exiting main")

}