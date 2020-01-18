package mypkg

// Person is a struct that has name and age.
type Person struct {
	Name string
	Age  int
}

// SayHello returns "Hello!" with Person's name.
func (p *Person) SayHello() string {
	return "Hello, I'm " + p.Name
}

// SayHello returns Person's age.
func (p *Person) SayAge() string {
	return "Hello, I'm " + string(p.Age) + " years old."
}
