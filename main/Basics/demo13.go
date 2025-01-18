package main

import "fmt"

type user1 struct {
	name     string
	password string
}

func (u user1) checkPassword(password string) bool {
	return u.password == password
}

func (u *user1) resetPassword(password string) {
	u.password = password
}

func main() {
	a := user1{name: "wang", password: "1024"}
	a.resetPassword("2048")
	fmt.Println(a.checkPassword("2048"))
}
