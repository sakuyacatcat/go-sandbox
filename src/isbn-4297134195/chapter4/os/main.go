package main

import (
	"log"
	"os"
)

func main() {
	f, err := os.Create("src/isbn-4297134195/chapter4/os/file.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	err = os.Mkdir("src/isbn-4297134195/chapter4/os/dir", 0755)
	if err != nil {
		log.Print("Error creating directory: ", err)
	}
	defer os.RemoveAll("src/isbn-4297134195/chapter4/os/dir")
}
