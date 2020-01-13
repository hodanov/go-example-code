package main

import "fmt"

// Vertex is a struct that has 2 int.
type Vertex struct {
	X, Y int
}

// Plus is a function to do the addition.
func (v *Vertex) Plus() int {
	return v.X + v.Y
}

func (v *Vertex) String() string {
	return fmt.Sprintf("X is %v! Y is %v!", v.X, v.Y)
}

func main() {
	v := Vertex{3, 4}
	fmt.Println(v.Plus())
	fmt.Println(&v)
}
