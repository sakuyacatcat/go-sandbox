package main

import "fmt"

func main() {
	x := 2
	y := 3

	if x == 2 && y == 3 {
		fmt.Println("x is 2 and y is 3")
	}

	i := 5

	switch i {
	case 1:
		fmt.Println("i is 1")
	case 2:
		fmt.Println("i is 2")
	case 3, 4, 5:
		fmt.Println("i is 3 or 4 or 5")
	default:
		fmt.Println("i is not between 1 and 5")
	}
}
