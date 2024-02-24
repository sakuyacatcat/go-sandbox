package main

import "fmt"

type User struct {
	Name string
	Age int
}

type Value int

func (u *User) AddAge(n Value) {
	u.Age += int(n)
}

func main() {
	user := new(User)
	user.Name = "Taro"
	user.Age = 20
	user.AddAge(5)
	fmt.Println(user)
	fmt.Println(&user)
	fmt.Println(*user)
}
