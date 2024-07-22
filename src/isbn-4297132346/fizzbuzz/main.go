package main

import (
	"fmt"

	"github.com/sakuyacatcat/fizzbuzz/pkg/handler"
)

func main() {
	fmt.Println("Starting FizzBuzz")

	h := handler.NewHandler()
	h.Handle()

	fmt.Println("Ending FizzBuzz")
}
