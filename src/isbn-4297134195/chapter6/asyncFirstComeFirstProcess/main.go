package main

import (
	"fmt"
	"time"
)

func main() {
	ch := fanIn(generator("Hello"), generator("Bye"))
	for i := 0; i < 10; i++ {
		fmt.Println(<-ch)
	}
}

func fanIn(ch1, ch2 <-chan string) <-chan string {
	newCh := make(chan string)
	go func() {
		for {
			select {
			case s := <-ch1: newCh <- s
			case s := <-ch2: newCh <- s
			}
		}
	}()
	return newCh
}

func generator(msg string) <-chan string {
	ch := make(chan string)
	go func() {
		for i := 0; ; i++ {
			ch <- fmt.Sprintf("%s %d", msg, i)
			time.Sleep(time.Second)
		}
	}()
	return ch
}
