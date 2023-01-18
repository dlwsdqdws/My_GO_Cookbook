package main

import (
	"fmt"
	"sync"
)

func hello(i int) {
	fmt.Println("hello ", i)
}

func main() {
	var wg sync.WaitGroup
	wg.Add(5)

}
