package main

import (
	"fmt"
	"sync"
)

func main() {
	// demoChannels()
	// demoSelect()
	// showLoopOverChannel()
	// demoLoopOverChannel()

}

func demoLoopOverChannel() {
	// Pull multiple vals
	ch := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			ch <- i
		}
		close(ch)
	}()
	for msg := range ch {
		fmt.Println(msg)
	}
}

func showLoopOverChannel() {
	ch := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			ch <- i
		}
		close(ch)
	}()

	for msg := range ch {
		fmt.Println(msg) // is gonna pull one val out of the channel with every loop
	}
}

func demoSelect() {
	ch1, ch2 := make(chan string), make(chan string)

	// go func() {
	// 	ch1 <- "message to channel 1"
	// }()
	// go func() {
	// 	ch2 <- "message to channel 2"
	// }()

	// time.Sleep(100 * time.Millisecond)

	select {
	case msg := <-ch1:
		fmt.Println("Chan1:", msg)
	case msg := <-ch2:
		fmt.Println("Chan2:", msg)
	default:
		fmt.Println("Default printing")
	}
	// No messages passed in either channel
	// we hit the select statement and no message on eith.
	// Deadlock - the runtime is gonna isssue a panic
}

func demoChannels() {
	var wg sync.WaitGroup
	ch := make(chan string)
	wg.Add(1)
	go func() {
		ch <- "the message"
		// wg.Done()
	}()
	go func() {
		fmt.Println(<-ch)
		wg.Done()
	}()
	wg.Wait()
}
