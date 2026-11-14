package main

import (
	"fmt"
	"math/rand/v2"
	"time"
)

func processNum(numChain chan int ){

	for num := range numChain {
		fmt.Println("data process ",num)
		time.Sleep(time.Second)
	}
	fmt.Println("processNum is activated")
}

func main() {
	numChain := make (chan int )
	fmt.Println("main funce ")
	go processNum(numChain)
	for{
		numChain <- rand.IntN(100)
	}

	// time.Sleep(time.Second*2)




	// messageChan := make(chan string)

	// messageChan <- "hello world"

	// mesg:= <- messageChan
	// fmt.Println("revieved message we got ",mesg)
}