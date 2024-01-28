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

	// slice with len and cap
	c := make([]int, 0, 100)
	for i := 0; i < 200; i++ {
		c = append(c, i)
		if i % 10 == 0 {
			fmt.Println(len(c), cap(c))
		}
	}

	// slice exclude specific element
	d := []int{1, 2, 3, 4, 5}
	d = append(d[:2], d[3:]...)
}
