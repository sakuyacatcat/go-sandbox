package main

import "fmt"

func server(ch chan string) {
	defer close(ch)
	ch <- "one"
	ch <- "two"
	ch <- "three"
}

func main() {
	ch := make(chan string)
	go server(ch)

	for s := range ch {
		fmt.Println(s)
	}
}
