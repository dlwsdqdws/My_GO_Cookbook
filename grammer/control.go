package main

import (
	"fmt"
	"time"
)

func main() {
	if num := 0; num < 0 {
		fmt.Println("negative")
	} else if num > 0 {
		fmt.Println("positive")
	} else {
		fmt.Println(0)
	}

	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("morning")
	case t.Hour() < 18:
		fmt.Println("afternoon")
	default:
		fmt.Println("evening")
	}

	for j := 0; j < 10; j++ {
		fmt.Println(j)
	}

	i := 1
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}
}
