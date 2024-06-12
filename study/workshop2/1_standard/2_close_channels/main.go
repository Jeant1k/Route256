package main

import (
	"fmt"
)

func main() {
	chValues := make(chan int, 5)

	go func() {
		defer close(chValues)

		for i := 1; i <= 5; i++ {
			chValues <- i * i
		}
	}()

	// time.Sleep(3 * time.Second)
	// value, ok := <- chValues

	// Читаем данные из канала пока он открыт
	// или пока в нем есть данные
	for i := range chValues {
		fmt.Println(i)
	}
}
