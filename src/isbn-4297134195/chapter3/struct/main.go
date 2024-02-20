package main

import "fmt"

type User struct {
	Name string
	Age int
}

func main() {
	// basic struct usage
	var user User
	user.Name = "Bob"
	user.Age  = 20

	// print struct without pointer. The value is copied every time.
	fmt.Printf("%#v\n", user)

	// print struct with pointer. The value is not copied.
	fmt.Printf("%#v\n", &user)
}
