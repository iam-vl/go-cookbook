package main

import "fmt"

func main() {
	age := 32
	// ageP := &age
	var ageP *int
	ageP = &age
	fmt.Println(*ageP) // dereference a pointer
	// fmt.Println("Age:", age)
	// fmt.Println(getAdultYears(age))
}

func getAdultYears(age int) int {
	return age - 18
}
