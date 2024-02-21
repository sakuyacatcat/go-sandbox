package main

import "fmt"

type User struct {
	Name string
	Age int
}

func main() {
	user := new(User)
	user.Name = "Taro"
	user.Age = 20
	fmt.Println(user)
	fmt.Println(&user)
	fmt.Println(*user)
}
