package main

import (
	"log"

	"github.com/sakuyacatcat/go-sandbox/src/isbn-4297134195/chapter4/functionOptionsPattern/server"
)

func main() {
	svr := server.New("localhost", 8888)
	if err := svr.Start(); err != nil {
		log.Fatal(err)
	}
}
