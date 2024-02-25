package main

import (
	"fmt"
	"reflect"
)

func main() {
	// basic interface
	var v interface{}
	v = 1
	callInt(v.(int))

	// interface with type switch
	var a interface{}
	a = 1
	printDetails(a)
	a = "Hello"
	printDetails(a)
	a = 1.0
	printDetails(a)

	// interface with original type
	type V int
	var b V = 1
	printDetailsReflect(b)
}

func callInt(num int) {
	fmt.Println(num)
}

func printDetails(a interface{}) {
	switch v := a.(type) {
	case int:
		fmt.Println("Integer", v)
	case string:
		fmt.Println("String", v)
	default:
		fmt.Println("Unknown type")
	}
}

func printDetailsReflect(a interface{}) {
	switch reflect.TypeOf(a).Kind() {
	case reflect.Int:
		fmt.Println("Integer", a)
	case reflect.String:
		fmt.Println("String", a)
	default:
		fmt.Println("Unknown type")
	}
}
