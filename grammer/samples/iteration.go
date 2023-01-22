package main

import "fmt"

func main() {
	m := map[string]string{"a": "A", "b": "B"}
	for k, v := range m {
		fmt.Println(k, v)
	}
	for k := range m {
		fmt.Println("key", k)
	}
	for _, v := range m {
		fmt.Println("value", v)
	}

	nums := []int{2, 3, 4}
	for index, value := range nums {
		fmt.Println(index, value)
	}
}
