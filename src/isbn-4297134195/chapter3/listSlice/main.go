package main

import "fmt"

func main() {
	// list
	var a [4]bool
	a[0] = true

	// slice
	b := make([]int, 4)
	b[0] = 1

	// append
	b = append(b, 3)

	fmt.Println(a)
	fmt.Println(b)
}
