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
}

func sendMessage(message string) {
	fmt.Println(message)
}
