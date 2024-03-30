package main

import (
	"fmt"
)

func main() {
	type human struct {
		name string
		age int
	}
	s := &human{
		name: "Alice",
		age: 28,
	}

	fmt.Printf("%+v\n", s)
	fmt.Printf("%#v\n", s)
}
