package mypkg

import (
	"testing"
)

func TestAverage(t *testing.T) {
	tests := []struct {
		name  string
		input []int
		want  int
	}{
		{"All positive integers", []int{1, 2, 3, 4, 5}, 3},
		{"Partially negative integer", []int{-1, 2, 3, 4, 5}, 2},
		{"Valid negative integer", []int{1, 2, 3, -4, -5}, 0},
		{"All negative integers", []int{-1, -2, -3, -4, -5}, -3},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			output := Average(tt.input)
			if output != tt.want {
				t.Errorf("want %d; got %d; input %d", tt.want, output, tt.input)
			}
		})
	}
}
