package something_test

import (
	"fmt"
	"testing"
)

func FuzzDoSomething(f *testing.F) {
	f.Add("test&&&")
	f.Fuzz(func(t *testing.T, input string) {
		doSomething(input)
	})
}

func doSomething(input string) {
	fmt.Println(input)
}
