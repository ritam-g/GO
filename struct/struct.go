package main

import (
	"fmt"
	"time"
)

type order struct {
	id        string
	status    string
	amount    float64
	createdAt time.Time
}

//pointer matters becuse end of the we are passing like varible right so it should point with same pointer as well 
func (o *order) changeStatus(status string){
	o.status=status
}

// pointer not reqruied here becuse we are not updating the value we are just changing the value ok 
func (o order) getAmount()float64{
	return o.amount
}

func main() {

	order:= order{
		id: "1",
		status: "pending",
		amount: 100,
		createdAt: time.Now(),
		
	}
	order.createdAt= time.Now()
	order.changeStatus("completed")
	fmt.Println("Ur new order details ",order)
	fmt.Println("order amount is ",order.getAmount())

}