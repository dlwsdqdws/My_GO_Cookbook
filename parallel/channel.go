package main

import "fmt"

func powFunc() {
	src := make(chan int)
	dest := make(chan int, 3)

	// Producer sends numbers
	go func() {
		defer close(src) // Close after src ends to reduce resource waste

		for i := 0; i < 10; i++ {
			src <- i
		}
	}()

	// Consumer calculates the square of the input numbers
	go func() {
		defer close(dest)

		// If the src channel did not close, then the range function will not end and cause a deadlock!
		for i := range src {
			dest <- (i * i)
		}
	}()

	// Main goroutine outputs the final results
	for i := range dest {
		fmt.Println(i)
	}
}

func main() {
	powFunc()
}
