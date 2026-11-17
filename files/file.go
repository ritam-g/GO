
package main

import (
	"fmt"
	"os"
)

func main() {
	f, err := os.Open("example.txt")

	if err != nil {
		fmt.Println("Error occurred:", err)
		return
	}
	defer f.Close()

	fileInfo, err := f.Stat()

	if err != nil {
		fmt.Println("Error occurred:", err)
		return
	}

	fmt.Println("The file name is:", fileInfo.Name())
}
