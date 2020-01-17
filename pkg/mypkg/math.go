/*
Package mypkg is my special package.
*/
package mypkg

// Average returns the average of a series of numbers.
func Average(s []int) int {
	total := 0
	for _, i := range s {
		total += i
	}
	return int(total / len(s))
}
