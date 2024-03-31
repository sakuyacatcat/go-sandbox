package main

import (
	"fmt"
	"io"
	"log"
)

type Reader interface {
	Read(p []byte) (n int, err error)
}

type Writer interface {
	Write(p []byte) (n int, err error)
}

type ReadWriter interface {
	Reader
	Writer
}

type Foo struct {
	Data []byte
}

func (f *Foo) Write(p []byte) (n int, err error) {
	f.Data = append(f.Data, p...)
	return len(p), nil
}

func (f *Foo) Read(p []byte) (n int, err error) {
	if len(f.Data) == 0 {
		return 0, io.EOF
	}

	n = copy(p, f.Data)
	f.Data = f.Data[n:]
	return n, nil
}

func main() {
	foo := &Foo{}

	writeData := []byte("hello, world")
	n, err := foo.Write(writeData)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(n, "bytes written")

	readData := make([]byte, len(writeData))
	n, err = foo.Read(readData)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(n, "bytes read:", string(readData))
	fmt.Println("Data:", string(foo.Data))
}
