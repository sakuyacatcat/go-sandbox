package main

import (
	"fmt"
	"strconv"
)

func main() {
	// declare variable
	var l int
	var m int = 10
	n := 20
	const x = 10

	fmt.Println(l, m, n)

	// strings from two strings
	var s string = strconv.Itoa(m) + strconv.Itoa(n)
	fmt.Println(s)
}
