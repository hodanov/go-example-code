package main

import "fmt"

func main() {

	f := 1.11
	fmt.Println(int(f))

	s := []int{1, 2, 5, 6, 2, 3, 1}
	fmt.Println(s[2:4])

	m := map[string]int{
		"Mike":  20,
		"Nancy": 24,
		"Messi": 30,
	}
	fmt.Printf("%T %v\n", m, m)

	// =================
	// map
	// =================
	// m := map[string]int{
	// 	"apple":  100,
	// 	"banana": 200,
	// }
	// fmt.Println(m)
	// fmt.Println(m["apple"])
	// fmt.Println(m["banana"])

	// m["banana"] = 300
	// fmt.Println(m["banana"])

	// m["new"] = 500
	// fmt.Println(m)

	// fmt.Println(m["nothing"])

	// v, ok := m["apple"]
	// fmt.Println(v, ok)

	// v2, ok2 := m["nothing"]
	// fmt.Println(v2, ok2)

	// // makeは予めメモリ上に確保するので、からのmapを作って後から要素を追加できる。
	// m2 := make(map[string]int)
	// m2["pc"] = 500
	// fmt.Println(m2)

	// varで宣言した場合はメモリ上に容量を確保してないので、要素を追加できない。
	// var m3 map[string]int
	// m3["pc"] = 500
	// fmt.Println(m3)

	// =================
	// make, len, cap
	// =================
	// makeは予めメモリ上に値を確保する

	// var c []int
	// n := make([]int, 5)
	// fmt.Printf("len=%d, cap=%d, value=%v\n", len(n), cap(n), n)

	// c = make([]int, 0, 5)
	// fmt.Printf("len=%d, cap=%d, value=%v\n", len(c), cap(c), c)
	// for i := 0; i < 5; i++ {
	// 	c = append(c, i)
	// 	fmt.Println(c)
	// }
	// fmt.Println(c)
}
