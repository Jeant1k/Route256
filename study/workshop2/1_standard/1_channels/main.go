package main

import "fmt"

func main() {
	chInt := make(chan int)

	go func() {
		fmt.Println("Ждем данные из канала")
		v := <-chInt
		fmt.Println("Получили данные", v)
	}()

	chInt <- 1
}
