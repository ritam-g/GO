package main

import "fmt"

func main() {
	const langName = "go lang"
	const (
		personName= "ritam "
		age = 22
		role="backend developer"
	)
	fmt.Println(langName)
	fmt.Println("the person details" ,"name : ", personName,"age :",age,"role : ",role)
}