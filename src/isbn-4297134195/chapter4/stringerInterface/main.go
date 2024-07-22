package main

import "fmt"

type F struct {
	Name string
	Age  int
}

func (f *F) String() string {
	return fmt.Sprintf("%s: %d", f.Name, f.Age)
}

func main() {
	f := &F{
		Name: "Alice",
		Age:  28,
	}

	fmt.Printf("%v\n", f)
}
