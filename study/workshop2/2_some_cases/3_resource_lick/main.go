package main

import (
	"sync"
	"time"
)

func main() {
	wg := sync.WaitGroup{}
	timer := time.NewTicker(2 * time.Second)
	for {
		select {
		case <-timer.C:
			refresh(wg)
		}
	}
	wg.Wait()
}

func refresh(wg sync.WaitGroup) {
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			wg.Done()
			time.Sleep(10 * time.Minute)
		}()
	}
}
