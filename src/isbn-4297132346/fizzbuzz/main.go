package main

import (
	"fmt"

	"github.com/sakuyacatcat/fizzbuzz/pkg/handler"
)

func main() {
	fmt.Println("#####################")
	fmt.Println("# Starting FizzBuzz #")
	fmt.Println("#####################")

	h := handler.NewHandler()
	h.Handle()

	fmt.Println("#####################")
	fmt.Println("#  Ending FizzBuzz  #")
	fmt.Println("#####################")
}
