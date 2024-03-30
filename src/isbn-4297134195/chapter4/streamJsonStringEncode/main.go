package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type UserData struct {
	Age int `json:"Age"`
	Job string `json:"Job"`
	Skills []string `json:"Skills"`
}

type User struct {
	Name string `json:"Name"`
	Data UserData `json:"Data"`
}

func main() {
	f, err := os.Open("src/isbn-4297134195/chapter4/streamJsonStringEncode/input.json")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	var user User
	err = json.NewDecoder(f).Decode(&user)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%#v\n", user)
}
