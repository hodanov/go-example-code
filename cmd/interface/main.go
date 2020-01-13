package main

import (
	"fmt"
)

// Human is a interface that have a func "say()".
type Human interface {
	Say() string
	DriveCar()
}

// Person is a struct that have a Name.
type Person struct {
	Name string
}

// Say is a func for Person to say something.
func (p *Person) Say() string {
	p.Name = "Mr." + p.Name
	fmt.Println(p.Name)
	return p.Name
}

// DriveCar is a func to drive a car.
// Human can drive a car but Dog can't.
func (p *Person) DriveCar() {
	fmt.Println(p.Name + " is driving a car.")
}

func main() {
	var hoge Human = &Person{"Hoge"}
	hoge.Say()
	hoge.DriveCar()
}
