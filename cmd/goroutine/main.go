package main

import (
	"fmt"
	//	"sync"
	//	"time"
)

// func goroutine(s []int, c chan int) {
// 	sum := 0
// 	for _, v := range s {
// 		sum += v
// 		c <- sum
// 	}
// 	close(c)
// }
//
// func main() {
// 	s := []int{1, 2, 3, 4, 5}
// 	c := make(chan int, len(s))
// 	go goroutine(s, c)
// 	for i := range c {
// 		fmt.Println(i)
// 	}
// }

/*
データのやりとりをする場合（returnとかで値を返したりとか）、channelを使う。
make(chan int)がキューみたいな感じになってる。
バッファを用意して、channelの数を指定したい場合はmake(chan int, 2)とかにする。
*/
func goroutine(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum
}

func main() {
	s := []int{1, 2, 3, 4, 5}
	c := make(chan int)
	go goroutine(s, c)
	go goroutine(s, c)
	x := <-c
	fmt.Println(x)
	y := <-c
	fmt.Println(y)
}

/*
goroutineは軽量のスレッド（並行処理）
goroutineの処理が完了しなくても、全体の処理が終わってしまうケースがある。
それを避けるためにsync.WaitGroupを使う。
Done()が返るまで待ってくれる。
*/

// func goroutine(s string, wg *sync.WaitGroup) {
// 	for i := 0; i < 5; i++ {
// 		time.Sleep(200 * time.Millisecond)
// 		fmt.Println(s)
// 	}
// 	wg.Done()
// }
//
// func normal(s string) {
// 	for i := 0; i < 5; i++ {
// 		time.Sleep(200 * time.Millisecond)
// 		fmt.Println(s)
// 	}
// }
//
// func main() {
// 	var wg sync.WaitGroup
// 	wg.Add(1)
// 	go goroutine("world", &wg)
// 	normal("hello")
// 	wg.Wait()
// }
