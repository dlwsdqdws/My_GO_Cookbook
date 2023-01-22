package main

import "fmt"

func add(a int, b int) int {
	return a + b
}

func mul(a, b int) int {
	return a * b
}

func exist(m map[string]int, k string) (v int, err bool) {
	v, err = m[k]
	return v, err
}

func increase(a int) {
	a += 1
}

func increase2(a *int) {
	*a += 1
}

func main() {

	fmt.Println(add(1, 2))
	fmt.Println(mul(3, 4))

	var m = map[string]int{
		"one": 1,
		"two": 2,
	}

	fmt.Println(exist(m, "one"))
	fmt.Println(exist(m, "zero"))

	num := 1
	increase(num)
	fmt.Println(num)
	increase2(&num)
	fmt.Println(num)
}
