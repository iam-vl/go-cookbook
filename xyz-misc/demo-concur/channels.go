package main

import (
	"fmt"
	"sync"
)

func DemoForLoop() {
	var wg sync.WaitGroup
	// orders are blocking each other
	var (
		receivedOrdersCh = make(chan Order)
		validOrderCh     = make(chan Order)
		invalidOrderCh   = make(chan InvalidOrder)
	)

	go receiveOrders(receivedOrdersCh)
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
func ExampleForLoop() {
	ch := make(chan string, 3)
	for _, v := range [...]string{"foo", "bar", "baz"} {
		ch <- v
	}
	close(ch)
	for msg := range ch {
		fmt.Println(msg)
	}
}

func DemoSelect() {
	var wg sync.WaitGroup
	// orders are blocking each other
	var (
		receivedOrdersCh = make(chan Order)
		validOrderCh     = make(chan Order)
		invalidOrderCh   = make(chan InvalidOrder)
	)

	go receiveOrders(receivedOrdersCh)
	go validateOrders(receivedOrdersCh, validOrderCh, invalidOrderCh)
	wg.Add(1) // orders are blocking each other

	go func(validOrderCh <-chan Order, invalidOrderCh <-chan InvalidOrder) {
		select {
		case order := <-validOrderCh:
			fmt.Printf("Valid order received: %v", order)
		case order := <-invalidOrderCh:
			fmt.Printf("Invalid order received: %v.Issue: %v\n", order.order, order.err)
		}
		wg.Done()
	}(validOrderCh, invalidOrderCh)

	// go func(validOrderCh <-chan Order) {
	// 	order := <-validOrderCh
	// 	fmt.Printf(`Valid order received: %v\n`, order)
	// 	wg.Done()
	// }(validOrderCh)

	// go func(invalidOrderCh <-chan InvalidOrder) {
	// 	order := <-invalidOrderCh

	// 	wg.Done()
	// }(invalidOrderCh)

	wg.Wait()
}

func BasicSelect() {
	ch1 := make(chan int, 1)
	ch2 := make(chan string, 1)

	// ch1 <- 999
	// ch2 <- "message"

	select {
	case msg := <-ch1:
		fmt.Println(msg)
	case msg := <-ch2:
		fmt.Println(msg)
	default:
		fmt.Println("Default")
	}
}

func DirChannelBasic() {
	ch := make(chan string)
	go func(ch chan<- string) {
		ch <- "message" // send-only channel
	}(ch)
	go func(ch <-chan string) {
		fmt.Println(<-ch)
	}(ch)
}

func DemoUnbufferedBuffered() {
	ch := make(chan string) // Unbuffered
	// ch := make(chan string, 1) // Buffered
	ch <- "message" // blocked for Unbuffered, no receiver. Works for buffered
	fmt.Println((<-ch))
}

func MainDemoOrders() {
	var wg sync.WaitGroup
	// orders are blocking each other
	var (
		receivedOrdersCh = make(chan Order)
		validOrderCh     = make(chan Order)
		invalidOrderCh   = make(chan InvalidOrder)
	)

	go receiveOrders(receivedOrdersCh)
	go validateOrders(receivedOrdersCh, validOrderCh, invalidOrderCh)
	wg.Add(1) // orders are blocking each other

	go func(validOrderCh <-chan Order) {
		order := <-validOrderCh
		fmt.Printf(`Valid order received: %v\n`, order)
		wg.Done()
	}(validOrderCh)

	go func(invalidOrderCh <-chan InvalidOrder) {
		order := <-invalidOrderCh
		fmt.Printf("Invalid order received: %v.Issue: %v\n", order.order, order.err)
		wg.Done()
	}(invalidOrderCh)

	wg.Wait()
}

// func receiveOrdersWg(wg *sync.WaitGroup) {
// 	for _, rawOrder := range rawOrders {
// 		var newOrder Order
// 		err := json.Unmarshal([]byte(rawOrder), &newOrder)
// 		if err != nil {
// 			log.Println(err)
// 			continue
// 		}
// 		orders = append(orders, newOrder)
// 	}
// 	wg.Done()
// }
