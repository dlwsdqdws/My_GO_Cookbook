package main

import (
	"fmt"
	"reflect"
)

type user struct {
	name string
	id   int
}

func (u user) checkId(id int) bool {
	return u.id == id
}

func main() {
	var user1 = user{}
	fmt.Println(user1)

	var user2 = user{name: "lulei"}
	fmt.Println(user2)

	var user3 = user{"lulei", 1}
	fmt.Println(user3)

	var user4 = user{}
	user4.id = 2
	fmt.Println(user4)

	var user5 = &user{}
	user5.name = "hahaha"
	fmt.Println(*user5)

	value := reflect.ValueOf(user3)
	for i := 0; i < value.NumField(); i++ {
		fmt.Println(i, value.Field(i))
	}

}
