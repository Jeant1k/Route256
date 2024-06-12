package main

import (
	"context"
	"fmt"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	var (
		ctx                   = context.Background()
		ctxWithCancel, cancel = signal.NotifyContext(ctx, syscall.SIGINT)
		chSync                = make(chan struct{})
	)
	defer cancel()

	go func(ctx context.Context, chSync chan<- struct{}) {
		for {
			select {
			case <-ctx.Done():
				fmt.Println("gracefull shutdown")
				time.Sleep(3 * time.Second)
				fmt.Println("Можем завершиться")
				chSync <- struct{}{}
				close(chSync)
				return
			default:
				fmt.Println("Я работаю!")
				time.Sleep(500 * time.Millisecond)
			}
		}
	}(ctxWithCancel, chSync)

	for range chSync {
		fmt.Println("Все горутины завершили работу")
	}
}
