package main

import (
	"errors"
	"fmt"
)

type user struct {
	name string
	id   int
}

func search(users []user, name string) (u *user, err error) {
	for _, u := range users {
		if u.name == name {
			return &u, nil
		}
	}

	return nil, errors.New("No such user")
}

func main() {
	u, err := search([]user{{"lulei", 1}, {"wong", 2}}, "leilu")
	if err != nil {
		panic(err)
	} else {
		fmt.Println(u.id)
	}
}
