package main

import (
	"fmt"
	"time"
)

func greet(phrase string, doneChannel chan bool) {
	fmt.Println("Hello!", phrase)
	doneChannel <- true
}

func slowGreet(phrase string, doneChannel chan bool) {
	time.Sleep(3 * time.Second)

	fmt.Println("Hello!", phrase)
	doneChannel <- true
	close(doneChannel)
}

// Concurrency
// func main() {
// 	greet("Nice to meet you!")
// 	greet("How are you?")
// 	slowGreet("How ... are ... you?")
// 	greet("I hope you're liking to learn go!!")
// }

// Parallelism with goroutines
func main() {
	// Channel to know if the goroutine has finished
	done := make(chan bool)
	// dones := make([]chan bool, 4)

	// dones[0] = make(chan bool)
	go greet("Nice to meet you!", done) // done would be dones[0]

	// dones[1] = make(chan bool)
	go greet("How are you?", done) // done would be dones[1]

	// dones[2] = make(chan bool)
	go slowGreet("How ... are ... you?", done) // done would be dones[2]

	// dones[3] = make(chan bool)
	go greet("I hope you're liking to learn go!!", done) // done would be dones[3]
 
	// Waiting for the channel to receive data(Go will wait until the channel receives data)
	for range done {
	}

	// for _, done := range dones {
	// 	<-done
	// }
}
