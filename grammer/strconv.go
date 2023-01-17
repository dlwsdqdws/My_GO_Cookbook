package main

import (
	"fmt"
	"strconv"
)

func main() {
	f, _ := strconv.ParseFloat("3.14159265", 64)
	fmt.Println(f)

	n, _ := strconv.ParseInt("1234", 10, 64)
	fmt.Println(n)

	n, _ = strconv.ParseInt("0x1001", 0, 64)
	fmt.Println(n)

	n2, _ := strconv.Atoi("123")
	fmt.Println(n2)

	n2, err := strconv.Atoi("A")
	fmt.Println(n2, err)
}
