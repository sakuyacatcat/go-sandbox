package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	go func() {
		for i := 0; i < 5; i++ {
			time.Sleep(1 * time.Second)
			ch1 <- "ch1-data"
		}
	}()

	go func() {
		for i := 0; i < 5; i++ {
			time.Sleep(2 * time.Second)
			ch2 <- "ch2-data"
		}
	}()

	for i := 0; i < 10; i++ {
		select {
		case msg1 := <-ch1:
			fmt.Println("ch1:", msg1)
		case msg2 := <-ch2:
			fmt.Println("ch2:", msg2)
		}
	}
}
