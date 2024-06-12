package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	chSignal := make(chan os.Signal)
	signal.Notify(chSignal, syscall.SIGINT)

	chSync := make(chan struct{})

	go func() {
		for {
			select {
			case <-chSync:
				fmt.Println("gracefull shutdown")
				time.Sleep(3 * time.Second)
				fmt.Println("Можем завершиться")
				chSync <- struct{}{}
			default:
				fmt.Println("Я работаю!")
				time.Sleep(500 * time.Millisecond)
			}
		}
	}()

	<-chSignal
	chSync <- struct{}{}
	<-chSync

}
