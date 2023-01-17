package main

import "fmt"

type user struct {
	name string
	id   int
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
}
