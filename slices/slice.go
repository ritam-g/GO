package main 

import "fmt"
func main (){
	var slice []int
	for i :=0;i<10;i++{
		slice = append(slice,i )
	}
	var slice1 =make([]bool,1)
	
	fmt.Println(slice)
	fmt.Println(slice==nil)
	fmt.Println(len(slice))

	//testing slice1
	fmt.Println(slice1)
	fmt.Println(slice1==nil)
	fmt.Println(len(slice1))


	//caps and len  
	var slice2 = make([]int,2)
	fmt.Println(slice2)
	fmt.Println(cap(slice2))
	fmt.Println(len(slice2))

	slice2=append(slice2,1,2,3,4,5)
	fmt.Println(slice2)
	fmt.Println(cap(slice2))
	fmt.Println(len(slice2))


}