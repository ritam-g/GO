package main

import ("fmt")

func main() {
	age := 8
	const role = "admin"
	var hasPermission bool = false
	if age >= 18 {
		fmt.Println("u are adult bro ")
	} else {
		fmt.Println("u are kid bro ")
	}

	if role == "admin" && hasPermission {
		fmt.Println("u are admin")
	} else {
		fmt.Println("u are not admin")
	}
	
}
