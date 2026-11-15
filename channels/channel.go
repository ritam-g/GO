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

func sendEmail(emailChain <-chan string, done chan <- bool){
	defer func(){done <- true}()

	for email:= range emailChain{
		fmt.Println("email sent ",email)
	}
}
func main() {

	chan1:= make(chan int)
	chan2:= make(chan string)


	go func(){chan1 <- 10}()
	go func(){chan2 <- "hellow"}()

	for i:=0;i<2;i++{
		select{
		case c1:= <-chan1:
			fmt.Println(c1)
		case c2:= <-chan2:
			fmt.Println(c2)
		}
	}

	// emailChain := make(chan string ,100)
	// done := make(chan bool)
	// for i:=range 100{
	// 	emailChain<- fmt.Sprintf("email %d",i)
	// }
	// // Tell the receiver: no more emails are coming
	// close(emailChain)

	// go sendEmail(emailChain,done)

	// <- done
	// emailChain <- "email 1"
	// emailChain <- "email 2"

	// fmt.Println(<-emailChain)
	// fmt.Println(<-emailChain)

	// done := make(chan bool)
	// go task(done)
	
	// <-done //kind of wait group blocking system 

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