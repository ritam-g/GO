package main

import (
	"fmt"
	"packages/auth"
	"packages/user"

	"github.com/fatih/color"
)

func main() {
	auth.LoginWithCredential("ritam ","maty")
	user := user.User{Name: "ritam", Age: 22}
	fmt.Println(user)
	color.Red(user.Name)
}