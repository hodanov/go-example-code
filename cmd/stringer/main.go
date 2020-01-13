package main

import "fmt"

// Person is a struct that has Name and Age.
type Person struct {
	Name string
	Age  int
}

func (p Person) String() string {
	return fmt.Sprintf("My name is %v.", p.Name)
}

func main() {
	mike := Person{"Mike", 22}
	// 通常はここで{Mike 22}と表示されるが、
	// String()で、出力の結果を変えている。
	fmt.Println(mike)
}
