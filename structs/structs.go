package main

import (
	"fmt"
	"time"
)

type user struct {
	name     string
	lastname string
	bithdate string
	created  time.Time
}

type Adnmin struct {
	email string
	user
}

func main() {
	user := user{
		name:     "Jhon",
		lastname: "Doe",
		bithdate: "MM/DD/YY",
		created:  time.Now(),
	}
	user.printUserFunctionReceiver()
	printUser(&user)
}

func (u *user) printUserFunctionReceiver() {
	fmt.Println(u)
}

func printUser(u *user) {
	fmt.Println(u)
}
