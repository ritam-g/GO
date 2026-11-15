package main

import (
	"fmt"
	"time"
)

func processNum(numChain chan int ){

	for num := range numChain {
		fmt.Println("data process ",num)
		time.Sleep(time.Second)
	}
	fmt.Println("processNum is activated")
}


// func sum(res chan int , n1 int ,n2 int){
// 	result:= n1+n2

// 	res <- result
// }

func task(done chan bool){
	defer func(){done <- true}()

	fmt.Println("processing....")
}
func main() {

	done := make(chan bool)
	go task(done)
	
	<-done //kind of wait group blocking system 

	// res := make(chan int )

	// go sum(res,10,20)

	
	// fmt.Println(<-res)

	// numChain := make (chan int )
	// fmt.Println("main funce ")
	// go processNum(numChain)
	// for{
	// 	numChain <- rand.IntN(100)
	// }

	// time.Sleep(time.Second*2)




	// messageChan := make(chan string)

	// messageChan <- "hello world"

	// mesg:= <- messageChan
	// fmt.Println("revieved message we got ",mesg)
}