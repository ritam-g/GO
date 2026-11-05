package main 

import "fmt"

func main(){
	// nums:= []int{1,2,3,4,5,6,7,8,9}
	// sum:=0
	// for i, num := range nums{
	// 	sum += num
	// 	fmt.Printf("index: %d, value: %d, sum: %d\n", i, num, sum)
	// }

	// fmt.Println("Total sum:", sum)


	// m:= map[string]string{
	// 	"name": "John",
	// 	"age": "30",
	// 	"city": "New York",
	// }

	// for k,v:= range m {
	// 	fmt.Printf("key: %s, value: %s\n", k, v)
	// }

	for k,v := range "golang"{
		fmt.Println(k,v)
		fmt.Printf("Value: %c\n", v)
	}
}