
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
		fmt.Println("working day", time.Now().Weekday())
	}

	whoAmI := func(i interface{}) {
		switch v := i.(type) {
		case int:
			fmt.Println("int:", v)

		case string:
			fmt.Println("string:", v)

		default:
			fmt.Println("unknown")
		}
	}

	whoAmI(10)
	whoAmI("Ritam")
	whoAmI(true)
	whoAmI(12.4)
}
