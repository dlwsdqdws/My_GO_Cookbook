package main

import "fmt"

func main() {
	s := make([]string, 3)
	s[1] = "lu"
	s[2] = "lei"
	fmt.Println(s)
	fmt.Println(len(s))

	s = append(s, "good")
	fmt.Println(s)
	fmt.Println(len(s))

	a := [5]int{1, 2, 3, 4, 5}
	var b []int = a[2:4]
	fmt.Println(b)
	fmt.Println(len(b))
	b = append(b, 6)
	fmt.Println(b)
	fmt.Println(len(b))

	good := []string{"g", "o", "o", "d"}
	fmt.Println(good)
	fmt.Println(len(good))
	good = append(good, ".")
	fmt.Println(good)
	fmt.Println(len(good))

	board := [][]int{
		[]int{1, 0, 1},
		[]int{0, 1, 1},
		[]int{1, 1, 0},
	}
	fmt.Println(board)
}
