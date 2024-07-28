package main

import "fmt"

func CreateIterator() {
	nextNum := sequenceGenerator()
	fmt.Println(nextNum())
	fmt.Println(nextNum())
}
func sequenceGenerator() func() int {
	i := 0
	return func() int {
		i += 1
		return i
	}
}
