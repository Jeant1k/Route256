package main

import (
	"fmt"
	"sync"
)

func main() {
	var (
		input      = make(chan int)
		out1, out2 = make(chan int), make(chan int)
		wg         sync.WaitGroup
	)

	wg.Add(2)
	go func() {
		defer wg.Done()
		for i := range input {
			out1 <- i * 2
		}

	}()
	go func() {
		defer wg.Done()
		for i := range input {
			out2 <- i
		}
	}()

	go func() {
		defer close(input)
		for i := 0; i < 5; i++ {
			input <- i
		}
	}()

	go func() {
		wg.Wait()
		close(out1)
		close(out2)
	}()

	go func() {
		for res := range out1 {
			fmt.Println(res)
		}
	}()

	for res := range out2 {
		fmt.Println(res)
	}

}
