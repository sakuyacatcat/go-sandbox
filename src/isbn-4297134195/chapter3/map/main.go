package main

import "fmt"

func main() {
	// basic map
	m := make(map[string]int)
	m["John"] = 21
	m["Bob"] = 18
	m["Mark"] = 33

	// map having cap
	mhc := make(map[string]int, 10)
	mhc["John"] = 21
	mhc["Bob"] = 18
	mhc["Mark"] = 33

	fmt.Println(mhc)

	// map iteration
	for k, v := range mhc {
		fmt.Printf("key: %v, value: %v\n", k, v)
	}
}
