package main

import (
	"context"
	"fmt"
	"golang.org/x/sync/semaphore"
	"time"
)

// semaphore...goroutineの数を調整できるパッケージ
var s *semaphore.Weighted = semaphore.NewWeighted(1)

func longProcess(ctx context.Context) {
	// //他のgoroutineを待たせる
	// if err := s.Acquire(ctx, 1); err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// defer s.Release(1)

	// 他のgoroutineが走ってたらs.Release(1)されるまでは何もせずにreturnする。
	isAcquire := s.TryAcquire(1)
	if !isAcquire {
		fmt.Println("Could not get lock")
		return
	}
	defer s.Release(1)

	fmt.Println("Wait...")
	time.Sleep(1 * time.Second)
	fmt.Println("Done")
}

func main() {
	ctx := context.TODO()
	go longProcess(ctx)
	go longProcess(ctx)
	go longProcess(ctx)
	time.Sleep(5 * time.Second)
}
