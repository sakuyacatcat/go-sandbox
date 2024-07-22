package main

import (
	"encoding/json"
	"fmt"
	"log"
)

var content = `{
	"Name": "Alice",
	"Data": {
		"Age": 28,
		"Job": "Engineer",
		"Skills": ["C", "Go", "Java"]
	}
}`

type UserData struct {
	Age    int      `json:"Age"`
	Job    string   `json:"Job"`
	Skills []string `json:"Skills"`
}

type User struct {
	Name string   `json:"Name"`
	Data UserData `json:"Data"`
}

func main() {
	var user User
	err := json.Unmarshal([]byte(content), &user)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%#v\n", user)
}
