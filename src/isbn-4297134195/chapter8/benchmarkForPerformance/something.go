package something

import "fmt"

func makeSomething(n int) []string {
	r := make([]string, n, n)
	for i := 0; i < n; i++ {
		r[i] = fmt.Sprintf("%05d 何か", i)
	}
	return r
}
