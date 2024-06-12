package main

import (
	"fmt"
)

// func main() {
// 	var (
// 		ch = make(chan struct{})
// 	)

// 	go func() {
// 		ch <- struct{}{}
// 	}()

// 	<-ch
// 	<-ch
// }

// func main() {
// 	var (
// 		wg = sync.WaitGroup{}
// 	)

// 	wg.Add(1)

// 	wg.Done()

// 	wg.Wait()
// }

func main() {
	var (
		ch = make(chan struct{})
	)

	go func() {
		ch <- struct{}{}
		// close(ch)
	}()

	for range ch {
		fmt.Println("bye bye")
	}
}
