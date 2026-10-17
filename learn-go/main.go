package main 

import "fmt"

func main(){
	if 20<10{
		fmt.Print("10 is lesser thatn 20 ")
	} else{
		fmt.Print("hey this is else part ")
	}

	hellow()
}


func hellow(){
	fmt.Println("hello")
	for i:=0;i<10;i++{
		fmt.Println(i)
	}
	fmt.Print(add(1,2))
}

func add(a int,b int) int {
	return a+b
}