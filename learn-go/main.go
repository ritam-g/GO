package main 

import (
	"fmt"
	"errors"
)

// func main(){
//     ans,err := devide(1,020)

// 	if err!=nil{
// 		fmt.Print("Error",err)
// 		return
// 	}

// 	fmt.Print(ans)
// }

func main(){
	defer fmt.Println("first")
	 fmt.Println("2nd")

	 prices := []int {1,2,3,4,5}
	 prices = append(prices,6)
	 fmt.Println(prices)
	 fmt.Println(len(prices))
	 fmt.Println(prices[3])
	fmt.Println("loop ")
	 for i ,v := range prices {
		fmt.Println(i,v)
	 }
}


func hellow(){
	fmt.Println("hello")
	for i:=0;i<10;i++{
		fmt.Println(i)
	}
	fmt.Println(add(1,2))
	fmt.Print(sub(1,2))
}

func add(a int,b int) int {
	return a+b
}

func sub(a,b int) int {
	return a-b
}


func devide(a,b float64) (float64,error){
	if b==0{
		return 0,errors.New("cannot devide by zero")
	}

	return a/b,nil
}