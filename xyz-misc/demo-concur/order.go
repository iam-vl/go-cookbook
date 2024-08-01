package main

import "fmt"

type Order struct {
	ProductCode int
	Quantity    float64
	Status      orderStatus
}

func (o Order) String() string {
	return fmt.Sprintf("Print code: %v, Qty: %v, Status: %v\n", o.ProductCode, o.Quantity, OrderStatusToText(o.Status))
}
func OrderStatusToText(o orderStatus) string {
	switch o {
	case none:
		return "none"
	case new:
		return "new"
	case received:
		return "received"
	case filled:
		return "filled"
	default:
		return "unknown status"
	}
}

type orderStatus int

var orders []Order

const (
	none orderStatus = iota
	new
	received
	reserved
	filled
)
