package main

import "fmt"

func main() {
	var a [5]int
	a[2] = 3
	fmt.Println(a)

	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println(b)

	var c [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			c[i][j] = (i + 1) * (j + 1)
		}
	}

	fmt.Println(c)
}
