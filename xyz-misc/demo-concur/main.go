package main

import "strconv"

var (
	in    = make(chan string)
	out   = make(chan int)
	errCh = make(chan error)
)

func worker(in <-chan string, out chan<- int, errChan chan<- error) {
	for msg := range in {
		val, err := strconv.Atoi(msg)
		if err != nil {
			errCh <- err
			return
		}
		out <- val
	}
}

func main() {

}

var rawOrders = []string{
	`{"ProductCode": 1111, "Quantity": -5, "Status": 1}`,
	`{"ProductCode": 2222, "Quantity": 42.3, "Status": 1}`,
	`{"ProductCode": 3333, "Quantity": 19, "Status": 1}`,
	`{"ProductCode": 4444, "Quantity": 8, "Status": 1}`,
}
