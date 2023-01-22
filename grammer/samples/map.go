package main

import "fmt"

func main() {
	m := make(map[string]int)

	m["one"] = 1
	m["two"] = 2

	fmt.Println(len(m))
	fmt.Println(m["one"])
	fmt.Println(m["zero"])

	v := m["two"]
	fmt.Println(v)
	v, ok := m["zero"]
	fmt.Println(v, ok)

	delete(m, "one")

	var m2 = map[string]int{
		"four": 4,
		"five": 5,
	}

	fmt.Println(m2)

	for k, v := range m2 {
		fmt.Println(k, v)
	}
}
