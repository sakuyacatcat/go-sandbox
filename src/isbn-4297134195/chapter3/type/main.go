package main

import "fmt"

// type declaration for primitive type
type MyString string

// type declaration for struct
type User struct {
	Name string
	Age int
}

func main() {
	// basic type
	var m MyString
	m = "foo"
	fmt.Println(m)
}
