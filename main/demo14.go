package main

import (
	"errors"
	"fmt"
)

type user2 struct {
	name     string
	password string
}

func findUser(users []user2, name string) (v *user2, err error) {
	for _, u := range users {
		if u.name == name {
			return &u, nil
		}
	}
	return nil, errors.New("not found")
}
func main() {
	u, err := findUser([]user2{{"wang", "1024"}}, "wang")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(u.name)

	if u, err := findUser([]user2{{"wang", "1024"}}, "li"); err != nil {
		fmt.Println(err)
		return
	} else {
		fmt.Println(u.name)
	}
}
