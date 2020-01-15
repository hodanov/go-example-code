package main

import (
	"testing"
)

func TestFibonacci(t *testing.T) {
	tests := []struct {
		name    string
		n, want int
	}{
		{"1st Fibonacci number", 0, 0},
		{"2nd Fibonacci number", 1, 1},
		{"10th Fibonacci number", 10, 55},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := fibonacci(tt.n)

			if got != tt.want {
				t.Errorf("want %d; got %d; input %d", tt.want, got, tt.n)
			}
		})
	}
}
