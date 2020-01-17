package mypkg

import (
	"fmt"
)

// Person is a struct that has name and age.
type Person struct {
	Name string
	Age  int
}

// Say returns "Hello!".
func (p *Person) Say() {
	fmt.Printf("Hello, I'm %v\n", p.Name)
}
