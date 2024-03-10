package main

import (
	"fmt"
	"log"
	"os"
	"path"
)

func main() {
	basePath := path.Dir("src/isbn-4297134195/chapter3/defer/main.go")
	f, err := os.Open(path.Join(basePath, "data.txt"))
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	printFile(f)

	f, err = os.Open(path.Join(basePath, "data1.txt"))
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	printFile(f)

	f, err = os.Open(path.Join(basePath, "data2.txt"))
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	printFile(f)
}

func printFile(f *os.File) {
	var b [512]byte
	n, err := f.Read(b[:])
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(b[:n]))
}
