package main

import (
	"context"
	"fmt"
	"sync"
)

// паттерн, который из множества каналов делает один канал
func main() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	n1, n2 := producer(1, 2, 3), producer(4, 5, 6)

	res := fanIn(ctx, n1, n2)
	for r := range res {
		fmt.Println(r)
	}
}

func fanIn(ctx context.Context, fetchers ...<-chan int) <-chan int {
	chCombine := make(chan int)

	var wg sync.WaitGroup
	wg.Add(len(fetchers))

	for _, f := range fetchers {
		go func(f <-chan int) {
			defer wg.Done()
			for {
				select {
				case res, ok := <-f:
					if !ok {
						return
					}
					chCombine <- res
				case <-ctx.Done():
					return
				}
			}
		}(f)
	}

	go func() {
		wg.Wait()
		close(chCombine)
	}()

	return chCombine
}

func producer(nums ...int) <-chan int {
	out := make(chan int)

	go func() {
		for _, n := range nums {
			out <- n
		}
		close(out)
	}()

	return out
}
