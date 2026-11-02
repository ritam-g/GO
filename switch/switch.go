package main

import (
	"fmt"
	"time"
)

func main() {
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("holiday")

	default:
		fmt.Println("working day",time.Now().Weekday())
	}
}
