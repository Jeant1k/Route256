package main

import (
	"fmt"
)

func main() {
	ch := make(chan struct{})

	go func() {
		for {
			// бесконечный цикл
		}
		ch <- struct{}{}
	}()

	<-ch // блокировка, но deadlock не сработает

	fmt.Println("Завершается...")
}
