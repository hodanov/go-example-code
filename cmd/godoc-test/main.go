package main

import (
	"app/pkg/mypkg"
	"fmt"
)

func main() {
	s := []int{1, 2, 3, 4, 5}
	fmt.Println(mypkg.Average(s))
}
