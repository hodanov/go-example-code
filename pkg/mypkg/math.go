/*
Package mypkg is my package.
*/
package mypkg

// Sum returns the sum of a series of numbers.
func Sum(s []int) int {
	total := 0
	for _, i := range s {
		total += i
	}
	return total
}
