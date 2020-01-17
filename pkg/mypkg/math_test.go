package mypkg

import (
	"fmt"
	"testing"
)

func Example() {
	v := Sum([]int{1, 2, 3, 4, 5})
	fmt.Println(v)
}

func ExampleSum() {
	v := Sum([]int{1, 2, 3, 4, 5})
	fmt.Println(v)
}

func TestSum(t *testing.T) {
	tests := []struct {
		name  string
		input []int
		want  int
	}{
		{"All positive integers", []int{1, 2, 3, 4, 5}, 15},
		{"Partially negative integer", []int{-1, 2, 3, 4, 5}, 13},
		{"Valid negative integer", []int{1, 2, 3, -4, -5}, -3},
		{"All negative integers", []int{-1, -2, -3, -4, -5}, -15},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			output := Sum(tt.input)
			if output != tt.want {
				t.Errorf("want %d; got %d; input %d", tt.want, output, tt.input)
			}
		})
	}
}
