package main

import (
	"fmt"
	"strings"
)

type user struct {
	name  string
	age   uint8
	email string
}

func (u user) String() string {
	return fmt.Sprintf("name:%v, age:%v, email:%v", u.name, u.age, u.email)
}

func main() {
	search_name := "mike"

	users := []user{
		{
			name:  "Mike",
			age:   32,
			email: "mike@gmail.com",
		},
		{
			name:  "John",
			age:   54,
			email: "john@gmail.com",
		},
		{
			name:  "Abobus",
			age:   2,
			email: "abobus@gmail.com",
		},
	}

	for _, user := range users {
		if strings.EqualFold(search_name, user.name) {
			fmt.Println(user)
		}
	}
}
