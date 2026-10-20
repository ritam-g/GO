package main

import (
	"fmt"
	"time"
)



func fetchPrice(symbol string,ch chan string){
	time.Sleep(1 * time.Second)

	ch <- "Price of " + symbol + " fetched"
}

func main(){
	// ch := make(chan string)

	// go func(){
	// 	ch <- "hello this is go cahnnel "
	// }()

	// msg := <-ch

	// fmt.Println(msg)

	symbols := []string{"INFY", "TCS", "WIPRO"}

	ch := make(chan string)

	for _,s:= range symbols{
		go fetchPrice(s,ch)
	}


	for i:=0; i <len(symbols);i++{
		fmt.Println(<-ch)
	}


}