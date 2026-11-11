package main

import (
	"fmt"
	"time"
)

func doTask(id int) {
	fmt.Println(id)
}

func main() {

	for i := range 10 {
		 go doTask(i)
	}

	time.Sleep(time.Second * 2)
}