package main

import (
	"bufio"
	"fmt"
	"os"
)

func countStrategy3(cources [][]float64, first, second, third int) float64 {
	var res float64
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			for k := 0; k < 3; k++ {
				if i != j && j != k && i != k {
					sum := cources[i][first] * cources[j][second] * cources[k][third]
					// fmt.Println(i, j, k, ":", sum)
					if sum > res {
						res = sum
					}
				}
			}
		}
	}
	// fmt.Println("strategy", first, second, third, ":", res)
	return res
}

func countStrategy2(cources [][]float64, first, second int) float64 {
	var res float64
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if i != j {
				sum := cources[i][first] * cources[j][second]
				// fmt.Println(i, j, ":", sum)
				if sum > res {
					res = sum
				}
			}
		}
	}
	// fmt.Println("strategy", first, second, ":", res)
	return res
}

func countStrategy1(cources [][]float64, first int) float64 {
	var res float64
	for i := 0; i < 3; i++ {
		sum := cources[i][first]
		// fmt.Println(i, ":", sum)
		if sum > res {
			res = sum
		}
	}
	// fmt.Println("strategy", first, ":", res)
	return res
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var t int
	fmt.Fscan(in, &t)

	for i := 0; i < t; i++ {
		cources := make([][]float64, 3)
		for j := 0; j < 3; j++ {
			cources[j] = make([]float64, 6)
			for k := 0; k < 6; k++ {
				var n, m float64
				fmt.Fscan(in, &n, &m)
				cources[j][k] = m / n
			}
		}

		// for _, bank := range cources {
		// 	fmt.Fprintln(out, strings.Trim(fmt.Sprint(bank), "[]"))
		// }

		results := []float64{
			countStrategy3(cources, 0, 2, 0),
			countStrategy3(cources, 0, 3, 5),
			countStrategy3(cources, 1, 4, 0),
			countStrategy2(cources, 1, 5),
			countStrategy1(cources, 0),
		}

		var max float64
		for _, result := range results {
			if result > max {
				max = result
			}
		}
		fmt.Fprintln(out, max)
	}
}
