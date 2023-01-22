package main

import (
	"fmt"
	"sync/atomic"
)

var x int32 = 100

func f_add() {
	atomic.AddInt32(&x, 1)
}

func f_sub() {
	atomic.AddInt32(&x, -1)
}

func main() {
	for i := 0; i < 100; i++ {
		f_add()
		f_sub()
	}
	fmt.Printf("x: %v\n", x)
}
