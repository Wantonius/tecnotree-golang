package main

import "fmt"

func willExecuteLast(greet string) {
	fmt.Printf("Goodbye %s, I was deferred to be last in the calling function\n",greet)
} 

func callsAdditionalDefer(greet string) {
	defer willExecuteLast(greet)
	fmt.Println("I will be before the first goodbye")
}

func helloGreeting(greet string) {
	fmt.Printf("Hello %s, I will execute first\n",greet)
}

//common use of panic is to abort if a function is about to return an
//error value and we do not want or can't handle it

func panics() {
	panic("calamity ensues!")
}

func main() {

	//recover must be called in a deferred function. After panic the 
	//deferred function will be called for recover

	defer func() {
		if r:=recover(); r != nil {
			fmt.Printf("It panicked but we recovered. Error:%s\n",r)
		}
	}()
	
	defer panics()
	
	defer fmt.Println("Next we panic and recover!")

	defer willExecuteLast("John")
	defer callsAdditionalDefer("Johnny")
	fmt.Println("First we test defer")
	helloGreeting("John")
}
