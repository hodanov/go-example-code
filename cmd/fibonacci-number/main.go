package main

import (
	"fmt"
)

func fibonacci(n int) int {
	if n == 0 || n == 1 {
		return n
	}
	return fibonacci(n-1) + fibonacci(n-2)
}

func main() {
	for i := 0; i < 30; i++ {
		fmt.Println(fibonacci(i))
	}
}
