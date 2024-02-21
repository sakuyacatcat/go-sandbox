package main

import "fmt"

type Value int

func (v Value) Add(n Value) Value {
	return v + n
}

func main() {
	var v Value = 1
	result := v.Add(2)
	fmt.Println(result)
}
