package main

import (
	"fmt"
	"time"
)

func main() {
	doneCh := make(chan bool)
	go greet("Nice to meet you!", doneCh)
	go greet("How are you?", doneCh)
	go slowGreet("How ... are ... you ...?", doneCh)
	go greet("I hope you're liking the course!", doneCh)

	for isDone := range doneCh {
		fmt.Println(isDone)
	}

}
func greet(phrase string, doneChan chan bool) {
	fmt.Println("Hello!", phrase)
	doneChan <- true
}

func slowGreet(phrase string, doneChan chan bool) {
	time.Sleep(3 * time.Second)
	fmt.Println("Hello!", phrase)
	doneChan <- true
	close(doneChan)
}
