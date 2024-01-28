package main

import "fmt"

func main() {
	s := "hello"
	// Golang string type is immutable. So, executing reassigning variable.
	s += "world"
	fmt.Println(s)

	// bytes type is mutable.
	b := []byte(s)
	b[0] = 'H'
	s = string(b)
	fmt.Println(s)

	// rune is unicode character.
	hello := "こんにちわ世界"
	r := []rune(hello)
	r[4] = 'は'
	s = string(r)
	fmt.Println(s)
}
