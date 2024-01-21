package main

import "fmt"

type Fruit int

func main() {
	const (
		Apple Fruit = iota
		Orange
		Banana
	)

	fmt.Println(Apple)
	fmt.Println(Orange)
	fmt.Println(Banana)
}
