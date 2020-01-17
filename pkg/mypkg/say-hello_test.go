package mypkg

import (
	"testing"
)

func TestSay(t *testing.T) {
	p := Person{Name: "Mike", Age: 3}
	tests := []struct {
		name string
		want string
	}{
		{"", "Hello, I'm " + p.Name},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			output := p.Say()
			if output != tt.want {
				t.Errorf("want %s; got %s", tt.want, output)
			}
		})
	}

}
