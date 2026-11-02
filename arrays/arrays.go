package main

import "fmt"

func main() {
	var arr1 [4]bool
	var arr2 [4]string
	var arr3 [4]int
	arr1[3] = true
	arr2[3] = "hey this is ritam "
	arr3[3] = 22

	fmt.Println(arr1)
	fmt.Println(arr2)
	fmt.Println(arr3)
}