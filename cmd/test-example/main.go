package main

import (
	"fmt"
)

func Calculate(x int) (result int) {
	result = x + 3
	return result
}

func main() {
	result := Calculate(3)
	fmt.Printf("3 + 2 = %v\n", result)
}
