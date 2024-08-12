package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"sync"
)

func mainDemoLoopReserve() {
	var wg sync.WaitGroup
	var (
		receivedOrdersCh             = receiveOrderEncapsulate()
		validOrderCh, invalidOrderCh = validateOrdersEncapsulate(receivedOrdersCh)
		reservedInventoryCh          = reserveInventory(validOrderCh) // new
	)
	// wg.Add(1)
	wg.Add(2)
	go func(invalidOrderCh <-chan InvalidOrder) {
		for order := range invalidOrderCh {
			fmt.Printf("Invalid order received: %v. Issue: %v\n", order.order, order.err)
		}
		wg.Done()
	}(invalidOrderCh)
	go func(reservedInventoryCh <-chan Order) {
		for order := range reservedInventoryCh {
			fmt.Printf("Inventory reserved for: %v\n", order)
		}
		wg.Done()
	}(reservedInventoryCh)
	wg.Wait()
}

func reserveInventory(in <-chan Order) <-chan Order {
	out := make(chan Order)
	go func() {
		for o := range in {
			o.Status = reserved
			out <- o
		}
		close(out)
	}()
	return out
}

func mainDemoLoopEncapsulate() {
	var wg sync.WaitGroup
	var (
		receivedOrdersCh             = receiveOrderEncapsulate()
		validOrderCh, invalidOrderCh = validateOrdersEncapsulate(receivedOrdersCh)
		// receivedOrdersCh = make(chan Order)
		// validOrderCh     = make(chan Order)
		// invalidOrderCh   = make(chan InvalidOrder)
	)
	wg.Add(1)
	// go receiveOrders(receivedOrdersCh) // No longer needed
	// go validateOrders(receivedOrdersCh, validOrderCh, invalidOrderCh)
	go func(validOrderCh <-chan Order, invalidOrderCh <-chan InvalidOrder) {
	loop:
		for {
			select {
			// ok - status of the channel
			case order, ok := <-validOrderCh:
				if ok {
					fmt.Println("Valid order received:", order)
				} else {
					break loop // break out of the loop
				}
			case order, ok := <-invalidOrderCh:
				if ok {
					fmt.Printf("Invalid order received: %v. Issue: %v\n", order.order, order.err)
				} else {
					break loop
				}
			}
		}
		wg.Done()
	}(validOrderCh, invalidOrderCh)

	wg.Wait()
}

// func validateOrdersEncapsulate(in <-chan Order, out chan<- Order, errCh chan<- InvalidOrder) {
// Receive only on the consumer side
func validateOrdersEncapsulate(in <-chan Order) (<-chan Order, <-chan InvalidOrder) {
	// Initialize channels
	var (
		out   = make(chan Order)
		errCh = make(chan InvalidOrder, 1) // only helps wi
	)
	// process in async
	go func() {
		for order := range in { // we can handle one error, can't handle multiple
			if order.Quantity > 0 {
				fmt.Println("Validated order: %v", order)
				out <- order
			} else {
				errCh <- InvalidOrder{order: order, err: errors.New("qty must be greater than zero")}
			}
		}
		close(out)
		close(errCh)
	}()
	return out, errCh
}
func receiveOrderEncapsulate() <-chan Order {
	out := make(chan Order) // new
	go func() {
		for _, rawOrder := range rawOrders {
			var newOrder Order
			err := json.Unmarshal([]byte(rawOrder), &newOrder)
			if err != nil {
				log.Println(err)
				continue
			}
			out <- newOrder
			// <-out // Compiler error
		}
		close(out)
	}()
	return out // new
}
func validateOrders(in <-chan Order, out chan<- Order, errCh chan<- InvalidOrder) {
	// order := <-in
	for order := range in {
		if order.Quantity > 0 {
			out <- order
		} else {
			errCh <- InvalidOrder{order: order, err: errors.New("qty must be greater than zero")}
		}
	}
	close(out)
	close(errCh)
}
func receiveOrders(out chan<- Order) {
	for _, rawOrder := range rawOrders {
		var newOrder Order
		err := json.Unmarshal([]byte(rawOrder), &newOrder)
		if err != nil {
			log.Println(err)
			continue
		}
		out <- newOrder
		// <-out // Compiler error
	}
	close(out)
}

var rawOrders = []string{
	`{"productCode": 1111, "quantity": -5, "status": 1}`,
	`{"productCode": 2222, "quantity": 42.3, "status": 1}`,
	`{"productCode": 3333, "quantity": 19, "status": 1}`,
	`{"productCode": 4444, "quantity": 8, "status": 1}`,
}

func mainDemoLoop01() {
	var wg sync.WaitGroup
	// New Channels
	var (
		receivedOrdersCh = make(chan Order)
		validOrderCh     = make(chan Order)
		invalidOrderCh   = make(chan InvalidOrder)
	)
	wg.Add(1)
	go receiveOrders(receivedOrdersCh)
	go validateOrders(receivedOrdersCh, validOrderCh, invalidOrderCh)
	go func(validOrderCh <-chan Order, invalidOrderCh <-chan InvalidOrder) {
	loop:
		for {
			select {
			// ok - status of the channel
			case order, ok := <-validOrderCh:
				if ok {
					fmt.Println("Valid order received:", order)
				} else {
					break loop // break out of the loop
				}
			case order, ok := <-invalidOrderCh:
				if ok {
					fmt.Printf("Valid order received: %v. Issue: %v\n", order.order, order.err)
				} else {
					break loop
				}
			}
		}
		wg.Done()
	}(validOrderCh, invalidOrderCh)

	wg.Wait()
}

func mainDemo01() {
	var wg sync.WaitGroup
	// New Channels
	var (
		receivedOrdersCh = make(chan Order)
		validOrderCh     = make(chan Order)
		invalidOrderCh   = make(chan InvalidOrder)
	)
	wg.Add(1)
	go receiveOrders(receivedOrdersCh)
	go validateOrders(receivedOrdersCh, validOrderCh, invalidOrderCh)
	go func(validOrderCh <-chan Order, invalidOrderCh <-chan InvalidOrder) {
		select {
		case order := <-validOrderCh:
			fmt.Println("Valid order received:", order)
		case order := <-invalidOrderCh:
			fmt.Printf("Valid order received: %v. Issue: %v\n", order.order, order.err)
		}
		wg.Done()
	}(validOrderCh, invalidOrderCh)

	wg.Wait()
}
