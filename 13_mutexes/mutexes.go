package main

import (
	"fmt"
	"sync"
)

var x = 0 //a global variable represents a critical section

func increment(wg *sync.WaitGroup) {
	x=x+1
	wg.Done()
}

func fixed_increment(wg *sync.WaitGroup, m *sync.Mutex) {
	m.Lock() //We lock the access to the critical section
	x=x+1
	m.Unlock() //Remember to unlock!
	wg.Done()	
}

func main() {
	var w sync.WaitGroup
	
	for i :=0;i<1000;i++ {
		w.Add(1)
		go increment(&w)
	}
	
	w.Wait()
	fmt.Println("Variable should be valued at 1000. It won't be in almost all cases. This is known as race condition")
	
	fmt.Printf("Final value of x is %d\n",x)

	fmt.Println("Lets fix this thing")
	x=0
	
	var m sync.Mutex

	for i :=0;i<1000;i++ {
		w.Add(1)
		go fixed_increment(&w,&m)
	}	
	
	w.Wait()
	
	fmt.Printf("Final value of x is %d\n",x)
}