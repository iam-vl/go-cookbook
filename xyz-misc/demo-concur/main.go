package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"strconv"
	"sync"
)

func main() {
	DemoPattern1()
}

func DemoPattern1() {
	var wg sync.WaitGroup
	// orders are blocking each other
	var (
		receivedOrdersCh = receiveOrders()
		validOrderCh     = make(chan Order)
		invalidOrderCh   = make(chan InvalidOrder)
	)

	// go receiveOrders(receivedOrdersCh)
	go validateOrders(receivedOrdersCh, validOrderCh, invalidOrderCh)
	wg.Add(1) // orders are blocking each other

	go func(validOrderCh <-chan Order, invalidOrderCh <-chan InvalidOrder) {
	loop:
		for {
			select {
			case order, ok := <-validOrderCh:
				if ok {
					fmt.Printf("Valid order received: %v", order)
				} else {
					break loop
				}

			case order, ok := <-invalidOrderCh:
				if ok {
					fmt.Printf("Invalid order received: %v.Issue: %v\n", order.order, order.err)
				} else {
					break loop
				}
			}
		}
		// select {
		// case order := <-validOrderCh:
		// 	fmt.Printf("Valid order received: %v", order)
		// case order := <-invalidOrderCh:
		// 	fmt.Printf("Invalid order received: %v.Issue: %v\n", order.order, order.err)
		// }
		wg.Done()
	}(validOrderCh, invalidOrderCh)

	wg.Wait()
}

// func validateOrders(in <-chan Order, out chan<- Order, errCh chan<- InvalidOrder) {
func validateOrders(in <-chan Order) (<-chan Order, <-chan InvalidOrder) {
	// order := <-in
	out := make(chan Order)
	errCh := make(chan InvalidOrder, 1)
	go func() {
		for order := range in {
			if order.Quantity <= 0 {
				errCh <- InvalidOrder{order: order, err: errors.New("Qty must be positive")}
			} else {
				out <- order
			}
		}
		close(out)   // Tell go we're done sending data
		close(errCh) // Tell go we're done sending data

	}()

	// if order.Quantity <= 0 {
	// 	errCh <- InvalidOrder{order: order, err: errors.New("Qty must be positive")}
	// } else {
	// 	out <- order
	// }
}
func receiveOrders() chan Order {
	// func receiveOrders(out chan<- Order) { // send-only channel
	out := make(chan Order)
	go func() {
		for _, rawOrder := range rawOrders {
			var newOrder Order
			err := json.Unmarshal([]byte(rawOrder), &newOrder)
			if err != nil {
				log.Println(err)
				continue
			}
			out <- newOrder
		}
		close(out) // tell go we're done sending
	}()
	return out
}

// func m2() {
// 	var (
// 		in = make(chan string)
// 	)
// 	out, errCh := worker2(in)

// }

func worker2(in <-chan string) (chan int, chan error) {
	out := make(chan int)
	errCh := make(chan error)
	go func() {
		for msg := range in {
			val, err := strconv.Atoi(msg)
			if err != nil {
				errCh <- err
				return
			}
			out <- val
		}
	}()
	return out, errCh
}

func m1() {
	var (
		in  = make(chan string)
		out = make(chan int)
		// What if errCh is never drained,
		// With a buff channel, errch is guaranteed to have a receiver, no leaks
		errCh = make(chan error, 1)
	)
	fmt.Println(in)
	fmt.Println(out)
	fmt.Println(errCh)

}

// func worker1(in <-chan string, out chan<- int, errChan chan<- error) {
// 	for msg := range in {
// 		val, err := strconv.Atoi(msg)
// 		if err != nil {
// 			errCh <- err
// 			return
// 		}
// 		out <- val
// 	}
// }

var rawOrders = []string{
	`{"ProductCode": 1111, "Quantity": -5, "Status": 1}`,
	`{"ProductCode": 2222, "Quantity": 42.3, "Status": 1}`,
	`{"ProductCode": 3333, "Quantity": 19, "Status": 1}`,
	`{"ProductCode": 4444, "Quantity": 8, "Status": 1}`,
}
