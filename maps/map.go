package main

import "fmt"

func main() {

	m := make(map[string]string)
	r:= make(map[string]int)

	m["name"] = "ritam maty "
	r["name"] =60
	r["class"] =0

	fmt.Println("the all roll values map ",r,len(r))
	// fmt.Printf("value of name is %v\n",m["name"])
	// fmt.Printf("value of roll is %v\n",r["ritam"])
	// fmt.Println("the length of r is ",len(r))
	delete(r,"class")

	fmt.Println("the all roll values map ",r,len(r))


	// remove all data 

	clear(r)
	fmt.Println("the all roll value map now ",r,len(r))
	
}