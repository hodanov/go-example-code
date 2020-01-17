package mypkg

// Person is a struct that has name and age.
type Person struct {
	Name string
	Age  int
}

// Say returns "Hello!".
func (p *Person) Say() string {
	return "Hello, I'm " + p.Name
}
