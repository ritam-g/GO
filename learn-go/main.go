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
    ch := make(chan int)

    go func() {
        for i := 1; i <= 5; i++ {
            ch <- i
        }
        close(ch) // done sending
    }()

    for v := range ch { // loops until channel is closed
        fmt.Println(v)
    }
}