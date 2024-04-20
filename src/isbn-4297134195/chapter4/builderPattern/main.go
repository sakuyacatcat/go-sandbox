package main

import (
	"log"
	"os"
	"time"

	"github.com/sakuyacatcat/go-sandbox/src/isbn-4297134195/chapter4/builderPattern/server"
)

func main() {
	f, err := os.Create("server.log")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	logger := log.New(f, "", log.LstdFlags)
	svr := server.NewBuilder("localhost", 8888).
		Timeout(time.Minute).
		Logger(logger).
		Build()
	if err := svr.Start(); err != nil {
		log.Fatal(err)
	}
}
