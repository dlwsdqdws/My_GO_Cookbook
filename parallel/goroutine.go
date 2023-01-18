package main

import (
	"fmt"
	"time"
)

func print(i int) {
	fmt.Println("hello goroutine : ", i)
}

func hello() {
	for i := 0; i < 5; i++ {
		go func(j int) {
			print(j)
		}(i)
	}

	// Create goroutine takes some time
	// Ensure that the main goroutine does not exit before the sub-goroutine is executed
	time.Sleep(time.Second)
}

func main() {
	hello()
}
