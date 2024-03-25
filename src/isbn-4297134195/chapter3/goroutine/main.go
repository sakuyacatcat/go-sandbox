package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	message := "hi"
	wg.Add(1)
	go func() {
		defer wg.Done()
		sendMessage(message)
	}()
	message = "bye"
	wg.Wait()
	fmt.Println("end of main")

	fmt.Println("Start race")
	race()
	fmt.Println("End race")
}

func sendMessage(message string) {
	fmt.Println(message)
}

func race() {
	n := 0

	var mu sync.Mutex

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		for i := 0; i<10000; i++ {
			mu.Lock()
			n++
			mu.Unlock()
		}
	}()

	go func() {
		defer wg.Done()
		for i := 0; i<10000; i++ {
			mu.Lock()
			n++
			mu.Unlock()
		}
	}()

	wg.Wait()

	fmt.Println(n)
}
