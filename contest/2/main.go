package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func isValidPin(pin string, letterMap map[string]int) bool {
	testMap := make(map[string]int)
	for _, char := range pin {
		testMap[string(char)]++
	}

	for letter, count := range letterMap {
		if testMap[letter] != count {
			return false
		}
	}
	return true
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	var n, t int
	fmt.Sscanf(scanner.Text(), "%d %d", &n, &t)

	scanner.Scan()
	letters := strings.Fields(scanner.Text())
	letterMap := make(map[string]int)
	for _, letter := range letters {
		letterMap[letter]++
	}

	for i := 0; i < t; i++ {
		scanner.Scan()
		testPin := scanner.Text()
		if isValidPin(testPin, letterMap) {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
	}
}
