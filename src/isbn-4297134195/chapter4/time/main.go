package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()
	fmt.Println(t.Format(time.RFC3339))

	ts := "2024/04/01 00:35:25"
	tt, err := time.Parse("2006/01/02 15:04:05", ts)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(tt.Format(time.RFC3339))

	td, err := time.ParseDuration("1h30m")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(td)
	fmt.Println(td * 3)
}
