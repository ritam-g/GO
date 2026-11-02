package main

import "fmt"

func main() {
	var arr1 [4]bool
	var arr2 [4]string
	var arr3 [4]int
	arr1[3] = true
	arr2[3] = "hey this is ritam "
	arr3[3] = 22

	arr4:= [7]int{1,2,3,4,5,6,7}
	arr5:=[2][3]int {
		{1,2,3},
		{4,5,6},
	}

	arr5[0][0]=10
	fmt.Println(arr4)
	fmt.Println(arr5)

	fmt.Println(arr1)
	fmt.Println(arr2)
	fmt.Println(arr3)
}