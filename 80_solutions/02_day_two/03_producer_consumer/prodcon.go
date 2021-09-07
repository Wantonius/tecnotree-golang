package main

import (
	"fmt"
	"time"
)

var messages = []string{
"The Waste Land by T.S.Eliot",
"April is the cruellest month, breeding",
"Lilacs out of the dead land, mixing",
"Memory and desire, stirring",
"Dull roots with spring rain.",
"Winter kept us warm, covering",
"Earth in forgetful snow, feeding",
"A little life with dried tubers.",
"Summer surprised us, coming over the Starnbergersee",
"With a shower of rain; we stopped in the colonnade,",
"And went on in sunlight, into the Hofgarten,",
"And drank coffee, and talked for an hour.",
"Bin gar keine Russin, stamm’ aus Litauen, echt deutsch.",
"And when we were children, staying at the arch-duke’s,",
"My cousin’s, he took me out on a sled,",
"And I was frightened. He said, Marie,",
"Marie, hold on tight. And down we went.",
"In the mountains, there you feel free.",
"I read, much of the night, and go south in the winter.",
}

const consumerCount int = 3

func produce(jobs chan<- string) {
	for _, msg := range messages {
		jobs <- msg
	}
	close(jobs)
}

func consume(worker int, jobs <-chan string, done chan<- bool) {
	for msg := range jobs {
		fmt.Printf("%v\n", msg)
		time.Sleep(2000*time.Millisecond)
	}
	done <- true
}

func main() {
	jobs := make(chan string)
	done := make(chan bool)

	go produce(jobs)

	for i := 1; i <= consumerCount; i++ {
		go consume(i, jobs, done)
	}
	<-done
} 
