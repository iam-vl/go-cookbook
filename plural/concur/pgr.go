package main

import (
	"fmt"
)

var ch = make(chan string, 1)

func main() {
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
