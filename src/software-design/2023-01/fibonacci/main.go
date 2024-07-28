package main

import "fmt"

var fibonacci = [100]int{}

func main() {
	num := 50
	fmt.Println(fibonacciByBottomUpDp(num))
	// fmt.Println(fibonacciByTopDownDp(num))
}

func fibonacciByTopDownDp(idx int) int {
	if idx == 0 {
		return 0
	} else if idx == 1 {
		return 1
	} else {
		if fibonacci[idx] == 0 {
			fibonacci[idx] = fibonacciByTopDownDp(idx-1) + fibonacciByTopDownDp(idx-2)
		}
		return fibonacci[idx]
	}
}

func fibonacciByBottomUpDp(idx int) int {
	fibonacci[0] = 0
	fibonacci[1] = 1
	for i := 2; i <= idx; i++ {
		fibonacci[i] = fibonacci[i-1] + fibonacci[i-2]
	}
	return fibonacci[idx]
}
