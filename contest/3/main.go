package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var t int
	fmt.Fscan(in, &t)

	for i := 0; i < t; i++ {
		var s string
		fmt.Fscan(in, &s)
		r := []rune(s)

		for i := 1; i < len(r)-1; i++ {
			if r[i-1] != r[i] && r[i] != r[i+1] {
				r = append(r[:i], r[i+1:]...)
			}
		}

		flag := true
		for i := 0; i < len(r)-1; i++ {
			if r[i] != r[i+1] {
				flag = false
				break
			}
		}

		if flag {
			fmt.Fprintln(out, "YES")
		} else {
			fmt.Fprintln(out, "NO")
		}
	}
}
