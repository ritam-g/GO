package main

import (
	"fmt"
	"time"
)



func fetchPrice(symbol string,ch chan string){
	time.Sleep(1 * time.Second)

	ch <- "Price of " + symbol + " fetched"
}

func main() {
    prices := make(chan string)
    orders := make(chan string)

    go func() {
        time.Sleep(1 * time.Second)
        prices <- "new price tick"
    }()

    go func() {
        time.Sleep(2 * time.Second)
        orders <- "order update"
    }()

    for i := 0; i < 2; i++ {
        select {
        case p := <-prices:
            fmt.Println("Got:", p)
        case o := <-orders:
            fmt.Println("Got:", o)
        }
    }
}