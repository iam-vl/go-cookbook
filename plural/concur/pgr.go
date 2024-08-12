package main

import (
	"fmt"
	"time"
)

func mainXFor() {
	ch := make(chan string, 3)
	for _, word := range [...]string{"foo", "bar", "baz"} {
		ch <- word
	}
	close(ch)
	for msg := range ch {
		fmt.Println(msg)
	}
}

func exampleSelect() {
	ch1 := make(chan int, 1)    // buffered chans prevent need for syncing
	ch2 := make(chan string, 1) // buffered chans prevent need for syncing
	ch1 <- 999
	ch2 <- "abc"
	select {
	case msg := <-ch1:
		fmt.Println(msg)
	case msg := <-ch2:
		fmt.Println(msg)
	}
}

func directional() {
	ch := make(chan string)
	go func(ch chan<- string) {
		ch <- "message"
	}(ch)
	go func(ch <-chan string) {
		fmt.Println(<-ch)
	}(ch)
	time.Sleep(time.Second * 1)
}
func bidir() {
	ch := make(chan string)
	go func(ch chan string) {
		ch <- "message"
	}(ch)
	go func(ch chan string) {
		fmt.Println(<-ch)
	}(ch)
	time.Sleep(time.Second * 1)
}

func unbuff() {
	ch := make(chan string, 1) // avoids deadlock - non-blocking operation
	ch <- "message"            // results in a deadlock with no receiver by default
	fmt.Println(<-ch)
}

var ch = make(chan string, 1)

func mainX01() {
	go sender()
	go receiver()
	// time.Sleep(1 * time.Second)
}

func sender() {
	ch <- "message"
}
func receiver() {
	msg := <-ch
	fmt.Println(msg)
}
