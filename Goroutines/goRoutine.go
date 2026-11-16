package main

import (
	"fmt"
	"sync"
)

func doTask(id int,wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println(id)
}

func main() {
	var wg sync.WaitGroup

	for i := range 10 {
		wg.Add(1)
		 go doTask(i,&wg)
	}

	wg.Wait()
}