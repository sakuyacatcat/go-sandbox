package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

func main() {
	for file := range FileIterator(".") {
		fmt.Println(file)
	}
}

func FileIterator(path string) <-chan string {
	out := make(chan string)
	go func() {
		defer close(out)
		files, err := ioutil.ReadDir(path)
		if err != nil {
			log.Fatal(err)
		}
		for _, file := range files {
			out <- file.Name()
		}
	}()
	return out
}
