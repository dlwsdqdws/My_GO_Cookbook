package main

import (
	"fmt"
	"strings"
)

type user struct {
	name string
	id   int
}

func main() {
	a := "hello"
	fmt.Println(strings.Contains(a, "ll"))
	fmt.Println(strings.Count(a, "l"))
	fmt.Println(strings.HasPrefix(a, "he"))
	fmt.Println(strings.HasSuffix(a, "llo"))
	fmt.Println(strings.Index(a, "ll"))
	fmt.Println(strings.Join([]string{"he", "llo"}, "-"))
	fmt.Println(strings.Repeat(a, 2))
	fmt.Println(strings.Replace(a, "e", "E", -1))
	fmt.Println(strings.Split("a-b-c", "-"))
	fmt.Println(strings.ToLower(a))
	fmt.Println(strings.ToUpper(a))
	fmt.Println(len(a))

	u := user{"lulei", 3}
	fmt.Printf("%v\n", u)
	fmt.Printf("%+v\n", u)
	fmt.Printf("%#v\n", u)

	f := 3.141592653
	fmt.Printf("%.2f\n", f)
}
