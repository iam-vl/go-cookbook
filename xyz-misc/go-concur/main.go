package main

import (
	"fmt"
	"time"
)

func greet(phrase string, doneChan chan bool) {
	fmt.Println("Hello!", phrase)
	// send data thru the channel once func is done/.
	doneChan <- true
}
func slowGreet(phrase string, doneChan chan bool) {
	time.Sleep(3 * time.Second)
	fmt.Println("Hello!", phrase)
	doneChan <- true
	// Explicitly close the channel
	close(doneChan)
}
func mainSlowGreet() {
	done := make(chan bool)

	go greet("Nice to meet you!", done)
	go greet("How are you?", done)
	go slowGreet("How ... are ... you ...?", done)
	go greet("I hope you're liking the course!", done)

	for range done {
		// fmt.Println(<-doneChan)
	}
	// for doneChan := range done {
	// 	fmt.Println(doneChan)
	// }

}

func main1() {
	doneChan := make(chan bool)
	// dones := make([]chan bool, 4)

	// dones[0] = make(chan bool)
	go greet("Nice to meet you!", doneChan)
	// dones[1] = make(chan bool)
	go greet("How are you?", doneChan)
	// dones[2] = make(chan bool)
	go slowGreet("How ... are ... you ...?", doneChan)
	// dones[3] = make(chan bool)
	go greet("I hope you're liking the course!", doneChan)

	for done := range doneChan {
		fmt.Println(done)
	}

	// <-doneChan // waiting for the data to come out of the channel
	// <-doneChan
	// <-doneChan
}
