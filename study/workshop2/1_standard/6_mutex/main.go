package main

import (
	"sync"
)

func main() {
	var (
		wg        = sync.WaitGroup{}
		mutex     = sync.Mutex{}
		sharedMap = map[int]int{}
	)

	wg.Add(2)
	go func() {
		defer wg.Done()
		for i := 0; i < 100; i++ {
			mutex.Lock()
			sharedMap[i] = i
			mutex.Unlock()
		}
	}()

	go func() {
		defer wg.Done()
		for i := 100; i < 200; i++ {
			mutex.Lock()
			sharedMap[i] = i
			mutex.Unlock()
		}
	}()

	wg.Wait()
}
