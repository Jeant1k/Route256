package main

import (
	"fmt"
	"sync"
	"time"
)

func processTask(taskNum int) {
	time.Sleep(1 * time.Second)
	fmt.Printf("Задача %d завершена\n", taskNum)
}

func main() {
	wg := sync.WaitGroup{}
	for i := 1; i <= 1000; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			processTask(i)
		}(i)
	}
	wg.Wait()
}
