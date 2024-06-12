package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	const (
		numJobs    = 5
		numWorkers = 5
	)

	var (
		jobs   = make(chan int, numJobs)
		result = make(chan int, numJobs)
		wg     = sync.WaitGroup{}
	)

	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go worker(i, &wg, jobs, result)
	}

	go func() {
		for i := 1; i <= numJobs; i++ {
			jobs <- i
		}
		close(jobs)
	}()

	// Wait в горутине, чтобы main не блокировался в ожидании
	go func() {
		wg.Wait()
		close(result)
	}()

	// вытаскивание данных до тех пор, пока данные есть, либо канал открыт
	for r := range result {
		fmt.Println("Результат:", r)
	}

}

func worker(id int, wg *sync.WaitGroup, jobs <-chan int, result chan int) {
	defer wg.Done()

	for job := range jobs {
		fmt.Printf("Воркер %d выполняет работу №%d\n", id, job)
		time.Sleep(2 * time.Second)
		result <- job * 2
	}
}
