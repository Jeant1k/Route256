package main

import (
	"fmt"
	"time"
)

func processTask(taskNum int) {
	time.Sleep(1 * time.Second)
	fmt.Printf("Задача %d завершена", taskNum)
}

func main() {
	for i := 0; i < 10; i++ {
		processTask(i)
	}
}
