package main

import (
	"context"
	"fmt"
)

func main() {
	data := make(chan interface{})
	go func() {
		for i := 0; i < 10; i++ {
			data <- i
		}
		close(data)
	}()

	chTakeFirstN := takeFirstN(context.Background(), data, 4)
	for v := range chTakeFirstN {
		fmt.Println(v)
	}
}

func takeFirstN(ctx context.Context, data <-chan interface{}, n int) <-chan interface{} {
	chTake := make(chan interface{})

	go func() {
		defer close(chTake)
		for i := 0; i < n; i++ {
			select {
			case val, ok := <-data:
				if !ok {
					return
				}
				chTake <- val
			case <-ctx.Done():
				return
			}
		}
	}()
	return chTake
}
