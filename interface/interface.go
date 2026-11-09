
package main

import "fmt"

type payment struct {
    gw paymneter
}

type paymneter interface {
    pay(amount float64)
    reFund(amount float64, reason string)
}

func (p payment) makePayment(amount float64) {
    p.gw.pay(amount)

    fmt.Println("payment gateway activated")
}

type rezorpay struct{}

func (r rezorpay) pay(amount float64) {
    fmt.Println("payment razorpay activated", amount)
}

type stripe struct{}

func (s stripe) pay(amount float64) {
    fmt.Println("payment stripe activated", amount)
}

type payPal struct{}

func (p payPal) pay(amount float64) {

     fmt.Println("payment razorpay activated", amount)
}
func (p payPal) reFund(amount float64, reason string) {

    fmt.Println("payment razorpay activated", amount, reason)
}

func main() {
    // razorpay := rezorpay{}
    // stripe := stripe{}
    payPal := payPal{}

    newPayment := payment{
        gw: payPal,
    }

    newPayment.makePayment(1000)
    newPayment.gw.reFund(1000, "due to location chagne ")
}
