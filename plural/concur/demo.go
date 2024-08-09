package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"sync"
)

func mainDemo01() {
	var wg sync.WaitGroup
	// New Channels
	var (
		receivedOrdersCh = make(chan Order)
		validOrderCh     = make(chan Order)
		invalidOrderCh   = make(chan InvalidOrder)
	)
	wg.Add(1)
	// go receiveOrders(&wg)
	go receiveOrders(receivedOrdersCh)
	go validateOrders(receivedOrdersCh, validOrderCh, invalidOrderCh)
	go func() {
		order := <-validOrderCh
		fmt.Println("Valid order received:", order)
		wg.Done()
	}()
	go func() {
		invalidOrder := <-invalidOrderCh
		fmt.Println("Invalid order received:", invalidOrder.order)
		wg.Done()
	}()
	wg.Wait()
	// fmt.Println(orders)
}

func validateOrders(in, out chan Order, errCh chan InvalidOrder) {
	// Grab order
	order := <-in
	if order.Quantity > 0 {
		out <- order
	} else {
		errCh <- InvalidOrder{order: order, err: errors.New("qty must be greater than zero")}
	}
}

// func receiveOrders(wg *sync.WaitGroup) {
func receiveOrders(out chan Order) {
	for _, rawOrder := range rawOrders {
		var newOrder Order
		err := json.Unmarshal([]byte(rawOrder), &newOrder)
		if err != nil {
			log.Println(err)
			continue
		}
		// orders = append(orders, newOrder)
		out <- newOrder
	}
	// wg.Done()
}

var rawOrders = []string{
	`{"productCode": 1111, "quantity": -5, "status": 1}`,
	`{"productCode": 2222, "quantity": 42.3, "status": 1}`,
	`{"productCode": 3333, "quantity": 19, "status": 1}`,
	`{"productCode": 4444, "quantity": 8, "status": 1}`,
}
