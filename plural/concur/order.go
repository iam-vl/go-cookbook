package main

import "fmt"

type Order struct {
	ProductCode int
	Quantity    float64
	Status      OrderStatus
}
type InvalidOrder struct {
	order Order
	err   error
}

func (o Order) String() string {
	return fmt.Sprintf("Prod code: %v, Qty: %v, Status: %v\n",
		o.ProductCode, o.Quantity, o.orderStatusToText())
}
func (o *Order) orderStatusToText() string {
	switch o.Status {
	case none:
		return "none"
	case new:
		return "new"
	case received:
		return "received"
	case reserved:
		return "reserved"
	case filled:
		return "filled"
	default:
		return "unknown status"
	}
}

type OrderStatus int

const (
	none OrderStatus = iota
	new
	received
	reserved
	filled
)
