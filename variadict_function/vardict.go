package main

import "fmt"

//veridct function which can take multiple arguments


func sum(nums ...int) int{
	total:=0

	for _,num:= range nums{
		total+=num
	}

	return total


}

func main (){
	fmt.Println(sum(1,2,3,4,5,6,7,8,8,10))
	nums:= []int{10,20,30,40,50}

	fmt.Println(sum(nums...))
}