package main

import "fmt"

type User struct {
	Name string
	Age int
}

func main() {
	var user User
	user.Name = "Bob"
	user.Age  = 20

	fmt.Printf("%#v\n", user)
}
