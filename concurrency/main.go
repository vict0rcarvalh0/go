package main

import (
	"errors"
	"fmt"
	"time"
)

func greet(phrase string, doneChannel chan bool, errorChan chan error) {
	_, err := fmt.Println("Hello!", phrase)
	doneChannel <- true

	if err != nil {
		errorChan <- err
		return
	}
}

func slowGreet(phrase string, doneChannel chan bool, errorChan chan error) {
	time.Sleep(3 * time.Second)

	_, err := fmt.Println("Hello!", phrase)

	errorChan <- errors.New("error found running Goroutine") // Forcing error

	if err != nil {
		errorChan <- err
		return
	}

	doneChannel <- true
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
	errorChan := make(chan error)

	go greet("Nice to meet you!", done, errorChan)
	go greet("How are you?", done, errorChan)
	go slowGreet("How ... are ... you?", done, errorChan)
	go greet("I hope you're liking to learn go!!", done, errorChan)

	// Waiting for the channel to receive data(Go will wait until the channel receives data)
	for i := 0; i < 4; i++ {
		select {
		case <-done:
			fmt.Println("Goroutine finished successfully")
		case err := <-errorChan:
			fmt.Println("Error found running Goroutine:", err)
		}
	}

	close(done)
	close(errorChan)
}
