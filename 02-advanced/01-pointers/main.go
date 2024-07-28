package main

import "fmt"

func main() {
	UpdateStructThruPointer()
}

func UpdateStructThruPointer() {
	p := Person{"John Doe", 30}
	p.BDay()
	fmt.Println(p)
	Birthday(&p)
	fmt.Println(p)
}

func (p *Person) BDay()  { p.Age += 1 }
func Birthday(p *Person) { p.Age += 1 }

type Person struct {
	Name string
	Age  int
}

func UpdatePrimitiveThruPinter() {
	var a int = 58
	var pA *int = &a
	fmt.Println("Address of a:", pA)
	// dereference p
	fmt.Println("Value of A thru pointer:", *pA)
	*pA = 123
	fmt.Println("New value of A:", a)
}
