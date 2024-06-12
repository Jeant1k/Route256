package main

import (
	"fmt"
	"bufio"
	"os"
	"strconv"
)

func main() {

	file, _ := os.Open("17.txt")
	defer file.Close()

	var numbers []int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		number, _ := strconv.Atoi(line)
		numbers = append(numbers, number)
	}

	count, max_sum := 0, 0
	for i := 0; i < len(numbers) - 1; i++ {
		if numbers[i] % 3 == 0 || numbers[i + 1] % 3 == 0 {
			count++;
			max_sum = max(max_sum, numbers[i] + numbers[i + 1])
		}
	}
	
	fmt.Println(count, max_sum)

}