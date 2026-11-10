package main

import "fmt"

func loops[T any](items []T) {
	fmt.Println("")
	for _, item := range items {
		fmt.Print(item)
	}
}
func loops2[T string| bool](items []T) {
	fmt.Println("")
	for _, item := range items {
		fmt.Print(item)
	}
}

type stack[T string | bool] struct{
	elements []T
}

func main() {


	var items = []int{1, 2, 3}
	fmt.Println("with genrics kind of all thing ")
	loops(items)

	var items2 = []string{"ritam", "maty"}
	var items3 = []bool{true, false}
	fmt.Println("with genrics which is string or bool ")
	loops2(items2)
	loops2(items3)

	ele:= stack[string]{
		elements: []string{"ritam","maty"},
	}

	fmt.Println(ele)
}