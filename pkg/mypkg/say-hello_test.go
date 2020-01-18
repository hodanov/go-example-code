package mypkg

import (
	"testing"
)

func ExamplePerson_SayHello() {
	p := Person{"Hoge", 13}
	p.SayHello()
}

func ExamplePerson_SayAge() {
	p := Person{"Hoge", 13}
	p.SayAge()
}

func TestSayHello(t *testing.T) {
	p := Person{Name: "Mike", Age: 3}
	tests := []struct {
		name string
		want string
	}{
		{"Say person's name", "Hello, I'm " + p.Name},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			output := p.SayHello()
			if output != tt.want {
				t.Errorf("want %s; got %s", tt.want, output)
			}
		})
	}
}

func TestSayAge(t *testing.T) {
	p := Person{Name: "Mike", Age: 3}
	tests := []struct {
		name string
		want string
	}{
		{"Say person's age", "Hello, I'm " + string(p.Age) + " years old."},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			output := p.SayAge()
			if output != tt.want {
				t.Errorf("want %s; got %s", tt.want, output)
			}
		})
	}
}
