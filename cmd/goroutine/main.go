package main

import (
	"fmt"
	"sync"
	//	"time"
)

/*
producerとconsumer
複数のサーバーのログを並列で取ってきて（producer）、
さらに同時並行で何か処理をさせたい（consumer）とかとか。
*/
func producer(ch chan int, i int) {
	ch <- i * 2
}

func consumer(ch chan int, wg *sync.WaitGroup) {
	for i := range ch {
		func() {
			defer wg.Done()
			fmt.Println("process", i*1000)
		}() // inner function
	}
}

func main() {
	var wg sync.WaitGroup
	ch := make(chan int)

	// Producer
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go producer(ch, i)
	}

	// Consumer
	go consumer(ch, &wg)
	wg.Wait()
	close(ch)
	fmt.Println("Done")
}

/*
channelとrangeとclose
for文で回して合計の計算途中の結果を逐次出力させたい場合。
*/
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
Buffered Channels
*/
// func main() {
// 	ch := make(chan int, 2) // バッファを2と指定してメモリを確保
// 	ch <- 100
// 	fmt.Println(len(ch)) // 1が出力される
// 	ch <- 200
// 	fmt.Println(len(ch)) // 2が出力される
// 	// ch <- 300 // バッファは2しかないのでエラーになる
// 	// どうしてもch <- 300を実行したい場合、一度channelを取り出す。
// 	// x := <-ch
// 	// fmt.Println(x)
// 	close(ch) // こうやって閉じておかないと、rangeで捌けなくなる
//
// 	for c := range ch {
// 		fmt.Println(c)
// 	}
// }

/*
channel
データのやりとりをする場合（returnとかで値を返したりとか）、channelを使う。
make(chan int)がキューみたいな感じになってる。
バッファを用意して、channelの数を指定したい場合はmake(chan int, 2)とかにする。
*/
// func goroutine(s []int, c chan int) {
// 	sum := 0
// 	for _, v := range s {
// 		sum += v
// 	}
// 	c <- sum
// }
//
// func main() {
// 	s := []int{1, 2, 3, 4, 5}
// 	c := make(chan int)
// 	go goroutine(s, c)
// 	go goroutine(s, c)
// 	x := <-c
// 	fmt.Println(x)
// 	y := <-c
// 	fmt.Println(y)
// }

/*
goroutineとsync.WaitGroup
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
