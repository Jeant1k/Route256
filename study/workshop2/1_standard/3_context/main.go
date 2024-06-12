package main

import (
	"context"
	"fmt"
	"time"
)

type myKey int

func main() {
	ctx := context.Background()
	ctxWithCancel, cancel := context.WithCancel(ctx)
	ctx = context.WithValue(ctx, myKey(1), 10)
	value := ctx.Value(myKey(1)).(int)

	go func() {
		fmt.Println("Я очень долго работаю")
		time.Sleep(3 * time.Second)
		cancel()
	}()

	<-ctxWithCancel.Done()
	fmt.Println("Меня отменили:", value)
}
