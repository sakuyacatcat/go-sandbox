package main

import (
	"github.com/sakuyacatcat/go-sandbox/src/isbn-4297134195/chapter4/internalPackage/example/internal/server"
)

func main() {
	svr := server.New("localhost", 8888)
	if err := svr.Start(); err != nil {
		panic(err)
	}
}
