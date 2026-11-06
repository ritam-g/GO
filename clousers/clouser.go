package main 

import "fmt"

func increase() func() int{
	count:=0

	return func()int{
		count +=1
		return count
	}
}
func main (){
	i := increase()

	fmt.Println("first icrement",i())
	fmt.Println("secondrst icrement",i())
	fmt.Println("thirdirst icrement",i())
	fmt.Println("fourthfirst icrement",i())

}