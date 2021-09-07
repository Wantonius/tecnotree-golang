package main

import (
	"fmt"
	"time"
)

func worker(id int, jobs <- chan int, results chan <- int) {
	for j := range jobs {
		fmt.Printf("Worker %d started a job\n",id)
		time.Sleep(3*time.Second)
		fmt.Printf("Worker %d finished a job\n",id)
		results <- j * 2
	}
}

func main() {

	var result int
	const numJobs = 10
	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)
	
	for w := 1; w <= 3; w++ {
		go worker(w, jobs, results)
	}
	
	for i := 1; i <=numJobs; i++ {
		jobs <- i
	}
	close(jobs)
	
	for a := 1; a <= numJobs;a++ {
		result = <-results
		fmt.Printf("Main: job result %d\n",result)
	}
	
}
