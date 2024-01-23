package main

import "fmt"

func main() {
	for i := 0; i < 5; i++ {
		fmt.Println("basic loop")
	}

	j := 0
	for j < 3 {
		fmt.Println("while loop")
		j++
	}

	k := 0
	for {
		fmt.Println("infinite loop")
		k++
		if k == 3 {
			break
		}
	}

	list := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for i, v := range list {
		fmt.Println("index:", i, "value:", v)
	}

	mapping := map[int]string{1: "a", 2: "b", 3: "c"}
	for k, v := range mapping {
		fmt.Println("key:", k, "value:", v)
	}
}
