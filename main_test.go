package main

import (
	"testing"
)

// シンプルなテストコード
func TestCalculate(t *testing.T) {
	output := Calculate(2)
	if output != 4 {
		t.Errorf("Calculate(2) = %v, expected 4", output)
	}
}

// テーブル駆動テストの手法に沿ったコード
func TestTableCalculate(t *testing.T) {
	var tests = []struct {
		input    int
		expected int
	}{
		{2, 4},
		{-1, 1},
		{0, 2},
		{-5, -3},
		{999, 1001},
	}

	for _, test := range tests {
		output := Calculate(test.input)
		if output != test.expected {
			t.Errorf("Test Failed: input %v, got %v, want %v", test.input, output, test.expected)
		}
	}
}
