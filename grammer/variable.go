package main

import (
	"fmt"
	"math"
)

func main() {
	var a = "init"

	g := a + "hhh"

	fmt.Println(g)

	var b, c int = 1, 2

	fmt.Println(b, c)

	const h = 500000000
	const i = 3e20 / h

	fmt.Println(math.Sin(i))

}
