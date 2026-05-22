package hello_world

import (
	"strconv"
	"fmt"
)

type UserInformation struct {
	name, address string
	age int
}

func DisplayUser(name string, address string, age int) string {
	return "Hello " + name + " " + address + " age " + strconv.Itoa(age)
}

func HelloWorld() string {
	user := UserInformation{
		name: "Jane Doe",
		address: "Bintaro",
		age: 25,
	}

	displayUser := DisplayUser

	return displayUser(user.name, user.address, user.age)
}