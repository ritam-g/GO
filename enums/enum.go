package main

import "fmt"

// type orderStatus int

// const (
// 	Received orderStatus = iota
// 	Confirmed
// 	Prepared
// 	Deliverd
// )

type orderStatus string

const (
	Received orderStatus = "received"
	Confirmed orderStatus = "confirmed"
	Prepared orderStatus = "prepared"
	Deliverd orderStatus = "deliverd"
)

func changeOrderStatus(status orderStatus) {
	fmt.Println("order status ",status)
}

func main() {

	changeOrderStatus(Received)
}